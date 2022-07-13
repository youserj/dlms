package common_data_type

import (
	"bytes"
	"fmt"
)

func Version() string {
	return "0.0.4"
}

// convert int to ASN.1 format
func encode_length(len uint32) []byte {
	if len < 0x80 {
		return []byte{byte(len)}
	} else if len < 0x1_00 {
		return []byte{0x81, byte(len)}
	} else if len < 0x1_00_00 {
		return []byte{0x82, byte(len >> 8), byte(len & 0xff)}
	} else if len < 0x1_00_00_00 {
		return []byte{0x83, byte(len >> 16), byte((len >> 8) & 0xff), byte(len & 0xff)}
	} else {
		return []byte{0x84, byte(len >> 24), byte((len >> 16) & 0xff), byte((len >> 8) & 0xff), byte(len & 0xff)}
	}
}

func decode_length(buf *bytes.Buffer) (uint32, error) {
	var tmp byte
	var err error
	var length uint32
	tmp, err = buf.ReadByte()
	length = uint32(tmp)
	if err != nil {
		return 0, fmt.Errorf("in reading length [%w]", err)
	} else if length < 0x80 {
		return length, nil
	} else if length == 0x80 {
		return 0, nil
	} else {
		length &= 0b0111_1111
		switch length {
		case 1:
			{
				tmp, err = buf.ReadByte()
				if err != nil {
					return 0, fmt.Errorf("in reading extended length byte [%w]", err)
				} else {
					length = uint32(tmp)
					return length, err
				}
			}
		case 2, 3, 4:
			{
				tmp2 := make([]byte, length)
				var n int
				n, err = buf.Read(tmp2)
				if err != nil {
					return 0, fmt.Errorf("in reading extended length %d bytes [%w]", n, err)
				} else if n != int(length) {
					return 0, &LengthError{n, int(length)}
				} else {
					length = 0
					for i, value := range tmp2 {
						length <<= i * 8
						length += uint32(value)
					}
					return length, nil
				}
			}
		default:
			return 0, &UnsupLengthError{int(length)}
		}
	}
}

type CDT interface {
	TAG() byte
	Encode() (ret []byte)
	// ContentsLen() uint32
	Set(*bytes.Buffer) (err error)
}

// write CDT encode to byte-buffer
func CDTtoBuffer(c CDT, buf *bytes.Buffer) (n int, err error) {
	n, err = buf.Write(c.Encode())
	return
}

// read expected tag from buffer
func read_tag(c CDT, buf *bytes.Buffer) (err error) {
	var read_tag byte
	read_tag, err = buf.ReadByte()
	if err != nil {
		return
	} else if read_tag != c.TAG() {
		err = &TagError{read_tag, c.TAG()}
		return
	}
	return
}

//For *NullData | *Integer | *Unsigned | *Enum | *Long | *LongUnsigned...
type CDTWithoutLength interface{
	CDT
	SetContents(*bytes.Buffer)error
	Contents()[]byte
}

// return Encoding TODO: maybe delete by exist .Encode in every CDT
func Encode[T CDTWithoutLength](c T)[]byte{
	ret := []byte{c.TAG()}
	return append(ret, c.Contents()...)
}

// Set from byteBuffer to Without Length DLMS types
func Set(c CDTWithoutLength, buf *bytes.Buffer)error{
	err := read_tag(c, buf)
	if err != nil {
		return err
	} else {
		return c.SetContents(buf)
	}
}

// return elemnt instance by it tag
func get_element_constuctor(tag byte) (func() CDT, error) {
	switch tag {
	case 0:
		return func() CDT { return new(NullData) }, nil
	case 1:
		return func() CDT { return new(Array) }, nil
	// case 2: el = new(Structure)
	// todo more 3..14
	case 15:
		return func() CDT { return new(Integer) }, nil
	// todo more 16..
	default:
		return nil, fmt.Errorf("unknown tag %d", tag)
	}
}
