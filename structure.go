package common_data_type

import (
	"bytes"
	"fmt"
)

type Structure struct{}

func (Structure) TAG()byte{
	return 2
}

type ScalerUnit struct{
	Structure
	Scaler Integer
	Unit UnitEnum
}

func (t *ScalerUnit) Length()uint32{
	return 2
}

// Todo: Can use constant 
func (t *ScalerUnit) ContentsLen()uint32{
	return t.Scaler.ContentsLen() + t.Unit.ContentsLen() 
}

func (t *ScalerUnit) Contents()[]byte{
	return t.Encode()[2:]
}
	
func (t *ScalerUnit) Encode()[]byte{
	ret := make([]byte, 6)
	ret[0] = t.TAG()
	ret[1] = 2
	copy(ret[2:], t.Scaler.Encode())
	copy(ret[4:], t.Unit.Encode())
	return ret
}

func (t *ScalerUnit) Set(buf *bytes.Buffer)error{
	err := read_tag(t, buf)
	if err != nil{
		return err
	} else {
		var length uint32
		length, err = decode_length(buf)
		if err != nil{
			return err
		} else if length != t.Length(){
			return fmt.Errorf("got length of Struct (here name) %d, expect %d", length, t.Length())
		} else{
			err = t.Scaler.Set(buf)
			if err != nil{
				return err
			} else{
				err = t.Unit.Set(buf)
				if err != nil{
					return err
				} else{
					return nil
				}
			}
		}
	} 
}