package common_data_type

import (
	"fmt"
	"bytes"
)

func Version() string {
	return "0.0.3"
}

type CDT interface{
	TAG() byte
	Encode() (ret []byte)
}

// write CDT encode to byte-buffer
func CDTtoBuffer(c CDT, buf *bytes.Buffer)(n int, err error){
	n, err = buf.Write(c.Encode())
	return
}

// read expected tag from buffer
func read_tag(tag byte, buf bytes.Buffer)(n int, err error){
	var read_tag byte
	read_tag, err = buf.ReadByte()
	if err != nil{
		return
	}else if read_tag != tag{
		err = fmt.Errorf("got tag %d, excepted %d", read_tag, tag)
		return
	}
	return
}

