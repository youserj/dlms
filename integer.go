package common_data_type

import (
	"bytes"
)

type integer struct {
	contents byte
}

func (*integer) ContentsLen() uint32 {
	return 1
}

func (c *integer) Contents() []byte {
	return []byte{c.contents}
}

func (c *integer) SetContents(buf *bytes.Buffer)error{
	contents, err := buf.ReadByte()
	if err != nil {
		return &LengthError{0, 1}
	} else {
		c.contents = contents
		return nil
	}
}

type Integer struct {
	integer
}

func (*Integer) TAG() byte {
	return 15
}

func (c *Integer) Encode() []byte {
	return []byte{c.TAG(), c.contents}
}

func (c *Integer) Set(buf *bytes.Buffer) error {
	return Set(c, buf)
} 

func (c *Integer) SetFromInt8(value int8) {
	c.contents = byte(value)
}

func (c *Integer) SetFromByte(value byte) {
	c.contents = byte(value)
}

func (c Integer) To_int8() int8 {
	return int8(c.contents)
}

type Unsigned struct {
	integer
}

func (*Unsigned) TAG() byte {
	return 17
}

func (c *Unsigned) Encode() []byte {
	return []byte{c.TAG(), c.contents}
}

func (c *Unsigned) Set(buf *bytes.Buffer) error {
	return Set(c, buf)
} 

func (c *Unsigned) SetFromUInt8(value uint8) {
	c.contents = byte(value)
}

func (c *Unsigned) SetFromByte(value byte) {
	c.contents = byte(value)
}

func (c Unsigned) To_uint8() uint8 {
	return uint8(c.contents)
}

type Enum struct{
	integer
}

func (*Enum) TAG() byte {
	return 22
}

func (c *Enum) Encode() (ret []byte) {
	return []byte{c.TAG(), byte(c.contents)}
}

func (c *Enum) Set(buf *bytes.Buffer) error {
	return Set(c, buf)
}

func (c *Enum) SetFromByte(value byte){
	c.contents = value
}
