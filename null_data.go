package common_data_type

import (
	"bytes"
	"fmt"
)

type NullData struct{}

func (NullData) TAG()byte{
	return 0
}

func (NullData) Contents()[]byte{
	return make([]byte, 0)
}

func (c NullData) Encode()(ret []byte){
	ret = []byte{c.TAG()}
	return
}

func (c *NullData) Write(buf bytes.Buffer)(n int, err error){
	var tag byte
	tag, err = buf.ReadByte()
	if err != nil{
		return
	}else if tag != c.TAG(){
		err = fmt.Errorf("got tag %d, excepted %d", tag, c.TAG())
		return
	}
	return
}

// read encode from nulldate to byte-buffer
func (c *NullData) Read(buf bytes.Buffer)(n int, err error){
	n, err = buf.Write(c.Encode())
	return
}