package common_data_type

import (
	"bytes"
)

type Integer struct{
	value int8
}

func (Integer) TAG()byte{
	return 15
}

// contents length
func (Integer) Len()int{
	return 1
}

func (c Integer) Contents()(ret []byte){
	ret = []byte{byte(c.value)}
	return
}

func (c Integer) Encode()(ret []byte){
	ret = []byte{c.TAG(), byte(c.value)}
	return
}

func (c *Integer) SetFromBuffer(buf bytes.Buffer)(n int, err error){
	n, err = read_tag(c.TAG(), buf)
	if err != nil{
		data := make([]byte, c.Len())
		n, err = buf.Read(data)
		n++
		return
	}
	return
}

// read encode from nulldate to byte-buffer
func (c *Integer) WriteToBuffer(buf bytes.Buffer)(n int, err error){
	n, err = buf.Write(c.Encode())
	return
}