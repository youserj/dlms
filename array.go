package common_data_type

import (
	"bytes"
	"fmt"
)

type Array struct {
	values []CDT
}

func (Array) TAG() byte {
	return 1
}

func (t Array) Length() uint32 {
	return uint32(len(t.values))
}

func (t *Array) Clear() []CDT {
	ret := t.values
	t.values = nil
	return ret
}

func (t *Array) Remove(index uint32) error {
	if index >= t.Length() {
		return fmt.Errorf("index %d out of Array range", index)
	}
	tmp := make([]CDT, t.Length()-1)
	copy(tmp, t.values[:index])
	copy(tmp[index:], t.values[index+1:])
	t.values = tmp
	return nil
}

func (t *Array) Encode() (ret []byte) {
	ret = append(ret, 1)
	ret = append(ret, encode_length(t.Length())...)
	for _, val := range t.values {
		ret = append(ret, val.Encode()...)
	}
	return
}

// todo: make get type from first element
func (t *Array) Set(buf *bytes.Buffer) error {
	err := read_tag(t, buf)
	if err != nil{
		return fmt.Errorf("in reading tag [%w]", err)
	} else {
		var length uint32
		length, err = decode_length(buf)
		if err != nil {
			return fmt.Errorf("in decoding [%w]", err)
		} else {
			keep := t.Clear()
			if buf.Len() == 0 {
				return fmt.Errorf("not enouth byte in reading tag of type struct")
			} else {
				var creator func() CDT
				creator, err = get_element_constuctor(buf.Bytes()[0])
				if err != nil {
					t.values = keep
					return fmt.Errorf("in finding type array element [%w]", err)
				}else{ 
					for i := 0; i < int(length); i++ {
						el := creator()
						err = el.Set(buf)
						if err != nil {
							t.values = keep
							return fmt.Errorf("in reading array element[%d] [%w]", i, err)
						} else {
							t.values = append(t.values, el)
						}
					}
					return nil
				}
			}
		}
	}
}

func (t *Array) Append(el CDT) {
	t.values = append(t.values, el)
}

func (t *Array) restore(els []CDT){
	t.values = els
}

// For all structures containing Array
type CDTArray interface{
	CDT
	Append(CDT)
	Clear() []CDT
	restore([]CDT)
	GetElement()CDT
}

func SetToArray(c CDTArray, buf *bytes.Buffer) error {
	err := read_tag(c, buf)
	if err == nil {
		var length uint32
		length, err = decode_length(buf)
		if err != nil {
			return fmt.Errorf("in decoding [%w]", err)
		} else {
			keep := c.Clear()
			for i := 0; i < int(length); i++ {
				new_el := c.GetElement()
				err = new_el.Set(buf)
				if err != nil {
					c.restore(keep)
					return fmt.Errorf("in reading array element[%d] [%w]", i, err)
				}
				c.Append(new_el)
			}
		}
	}
	return err
}

type IntegerArray struct {
	Array
}

func (IntegerArray) GetElement()CDT{
	return new(Integer)
}

func (t *IntegerArray) Set(buf *bytes.Buffer) error {
	return SetToArray(t, buf)
}