package common_data_type

import "fmt"

type Enum byte

func (*Enum) TAG() byte {
	return 22
}

func (*Enum) ContentsLen() int {
	return 1
}

func (c Enum) Contents() (ret []byte) {
	ret = []byte{byte(c)}
	return
}

func (c *Enum) Encode() (ret []byte) {
	ret = []byte{c.TAG(), byte(*c)}
	return
}

func (c *Enum) Set(value []byte) (err error) {
	*c = Enum(value[0])
	return
}

// Todo: separate to other module
type Unit struct {
	Enum
}

func (c *Unit) Set(value []byte) (err error) {
	switch value[0]{
	case 58, 59:
		err = fmt.Errorf("%d is reserved", value[0])
	default: 
		*c = Unit{Enum(value[0])}
	}
	return
}