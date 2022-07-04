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
		return 0, fmt.Errorf("in reading length; %s", err)
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
					return 0, fmt.Errorf("in reading extended length byte; %s", err)
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
					return 0, fmt.Errorf("in reading extended length %d bytes; %s", n, err)
				} else if n != int(length) {
					return 0, fmt.Errorf("got %d buffer length, expected %d", n, length)
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
			return 0, fmt.Errorf("unsupported length: %d", length)
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
		err = fmt.Errorf("got tag %d, excepted %d", read_tag, c.TAG())
		return
	}
	return
}

// read expected tag and length from buffer
// func read_tag_and_length

// For Integer, Unsigned, Enum
type CDTOneByte interface {
	CDT
	SetFromByte(byte) error
}

// ForLong, UnsignedLong
type CDTTwoByte interface {
	CDT
	SetFromTwoBytes([]byte) error
}

func SetOneByte(c CDTOneByte, buf *bytes.Buffer) error {
	err := read_tag(c, buf)
	if err != nil {
		return err
	} else {
		var contents byte
		contents, err = buf.ReadByte()
		if err != nil {
			return fmt.Errorf("not enough value with tag: %d", c.TAG())
		} else {
			return c.SetFromByte(contents)
		}
	}
}

func SetTwoByte(c CDTTwoByte, buf *bytes.Buffer) error {
	err := read_tag(c, buf)
	if err != nil {
		return err
	} else {
		contents := make([]byte, 2)
		var n int
		n, err = buf.Read(contents)
		if err != nil {
			return fmt.Errorf("not enough value with tag: %d", c.TAG())
		} else if n != 2 {
			return fmt.Errorf("expect 2 bytes, got %d", n)
		} else {
			return c.SetFromTwoBytes(contents)
		}
	}
}
