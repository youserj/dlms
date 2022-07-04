package common_data_type

import (
	"bytes"
	"fmt"
)

type Enum struct{
	contents byte
}

func (*Enum) TAG() byte {
	return 22
}

func (*Enum) ContentsLen() uint32 {
	return 1
}

func (c *Enum) Contents() (ret []byte) {
	ret = []byte{c.contents}
	return
}

func (c *Enum) Encode() (ret []byte) {
	ret = []byte{c.TAG(), byte(c.contents)}
	return
}

func (c *Enum) Set(buf *bytes.Buffer) error {
	return SetOneByte(c, buf)
}

func (c *Enum) SetFromByte(value byte)(err error){
	c.contents = value
	return
}

// Todo: separate to other module
type UnitEnum struct {
	Enum
}

func (c *UnitEnum) SetFromByte(value byte) (err error) {
	switch value{
	case 58, 59:
		err = fmt.Errorf("%d is reserved", value)
	default: 
		c.contents = value
	}
	return
}