package common_data_type

import "bytes"

type long struct {
	contents [2]byte
}

func (*long) ContentsLen() uint32 {
	return 2
}

func (c *long) Contents() []byte {
	return []byte{c.contents[0], c.contents[1]}
}

func (c *long) SetContents(buf *bytes.Buffer) error {
	contents := make([]byte, 2)
	n, err := buf.Read(contents)
	if err != nil {
		return &LengthError{0, 2}
	} else if n != 2 {
		return &LengthError{n, 2}
	} else {
		c.contents[0] = contents[0]
		c.contents[1] = contents[1]
		return nil
	}
}

type Long struct {
	long
}

func (*Long) TAG() byte {
	return 16
}

func (c *Long) Encode() []byte {
	return []byte{c.TAG(), c.contents[0], c.contents[1]}
}

func (c *Long) Set(buf *bytes.Buffer) error {
	return Set(c, buf)
}

func (c *Long) SetFromInt16(value int16) {
	c.contents[0] = byte(value >> 8)
	c.contents[1] = byte(value & 0xff)
}

func (c *Long) ToInt16() int16 {
	return int16(c.contents[1]) + int16(c.contents[0])<<8
}

type LongUnsigned struct {
	long
}

func (*LongUnsigned) TAG() byte {
	return 18
}

func (c *LongUnsigned) Encode() []byte {
	return []byte{c.TAG(), c.contents[0], c.contents[1]}
}

func (c *LongUnsigned) Set(buf *bytes.Buffer) error {
	return Set(c, buf)
}

func (c *LongUnsigned) SetFromUInt16(value uint16) {
	c.contents[0] = byte(value >> 8)
	c.contents[1] = byte(value & 0xff)
}

func (c *LongUnsigned) ToUInt16() uint16 {
	return uint16(c.contents[1]) + uint16(c.contents[0])<<8
}