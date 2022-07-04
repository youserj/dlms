package common_data_type

import "bytes"

type Long struct {
	contents [2]byte
}

func (*Long) TAG() byte {
	return 16
}

func (*Long) ContentsLen() uint32 {
	return 2
}

func (c *Long) Contents() (ret []byte) {
	ret = []byte{c.contents[0], c.contents[1]}
	return
}

func (c *Long) Encode() []byte {
	return []byte{c.TAG(), c.contents[0], c.contents[1]}
}

func (c *Long) Set(buf *bytes.Buffer)error{
	return SetTwoByte(c, buf)
}

func (c *Long) SetFromInt16(value int16) (err error) {
	c.contents[0] = byte(value >> 8)
	c.contents[1] = byte(value & 0xff)
	return
}

func (c *Long) SetFromTwoBytes(value []byte)(err error){
	c.contents[0] = value[0]
	c.contents[1] = value[1]
	return
}

func (c *Long) Decode() int16 {
	return int16(c.contents[1]) + int16(c.contents[0])<<8
}