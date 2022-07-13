package common_data_type

import (
	"fmt"
)

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