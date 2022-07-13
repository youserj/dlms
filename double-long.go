package common_data_type

import "bytes"

type double_long struct {
	contents [4]byte
}

func (*double_long) ContentsLen() uint32 {
	return 4
}

func (c *double_long) Contents() []byte {
	return []byte{c.contents[0], c.contents[1], c.contents[2], c.contents[3]}
}

func (c *double_long) SetContents(buf *bytes.Buffer)error{
	contents := make([]byte, 4)
	n, err := buf.Read(contents)
	if err != nil {
		return &LengthError{0, 4}
	} else if n != 2 {
		return &LengthError{n, 4}
	} else {
		c.contents[0] = contents[0]
		c.contents[1] = contents[1]
		c.contents[2] = contents[2]
		c.contents[3] = contents[3]
		return nil
	}
}

type DoubleLong struct {
	double_long
}

func (*DoubleLong) TAG() byte {
	return 5
}

func (c *DoubleLong) Encode() []byte {
	return []byte{c.TAG(), c.contents[0], c.contents[1], c.contents[2], c.contents[3]}
}

func (c *DoubleLong) Set(buf *bytes.Buffer)error{
	return Set(c, buf)
}

func (c *DoubleLong) SetFromInt32(value int32){
	c.contents[0] = byte(value >> 24)
	c.contents[1] = byte((value >> 16) & 0xff)
	c.contents[2] = byte((value >> 8) & 0xff)
	c.contents[3] = byte(value & 0xff)
}

func (c *DoubleLong) ToInt32() int32 {
	return int32(c.contents[3]) + int32(c.contents[2])<<8 + int32(c.contents[1])<<16 + int32(c.contents[0])<<24
}

type DoubleLongUnsigned struct {
	double_long
}

func (*DoubleLongUnsigned) TAG() byte {
	return 6
}

func (c *DoubleLongUnsigned) Encode() []byte {
	return []byte{c.TAG(), c.contents[0], c.contents[1], c.contents[2], c.contents[3]}
}

func (c *DoubleLongUnsigned) Set(buf *bytes.Buffer)error{
	return Set(c, buf)
}

func (c *DoubleLongUnsigned) SetFromUInt32(value uint32){
	c.contents[0] = byte(value >> 24)
	c.contents[1] = byte((value >> 16) & 0xff)
	c.contents[2] = byte((value >> 8) & 0xff)
	c.contents[3] = byte(value & 0xff)
}

func (c *DoubleLongUnsigned) ToUInt32() uint32 {
	return uint32(c.contents[3]) + uint32(c.contents[2])<<8 + uint32(c.contents[1])<<16 + uint32(c.contents[0])<<24
}