package common_data_type

import (
	"bytes"
)

type Integer struct {
	contents byte
}

func (*Integer) TAG() byte {
	return 15
}

func (*Integer) ContentsLen() uint32 {
	return 1
}

func (c *Integer) Contents() []byte {
	return []byte{c.contents}
}

func (c *Integer) Encode() []byte {
	return []byte{c.TAG(), c.contents}
}

func (c *Integer) Set(buf *bytes.Buffer) error {
	return SetOneByte(c, buf)
} 

func (c *Integer) SetFromInt8(value int8) (err error) {
	c.contents = byte(value)
	return
}

func (c *Integer) SetFromByte(value byte) (err error) {
	c.contents = byte(value)
	return
}

func (c Integer) Decode() int8 {
	return int8(c.contents)
}
