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
	ContentsLen() int
	Set([]byte) (err error)
}

// write CDT encode to byte-buffer
func CDTtoBuffer(c CDT, buf *bytes.Buffer)(n int, err error){
	n, err = buf.Write(c.Encode())
	return
}

func BufferToCDT(c CDT, buf *bytes.Buffer)(n int, err error){
	n, err = read_tag(c.TAG(), buf)
	if err != nil{
		// error
	} else{
		data := make([]byte, c.ContentsLen())
		n, err = buf.Read(data)
		if err != nil{
			// error
		} else{
			n++
			c.Set(data)
		}
	}
	return
}

// read expected tag from buffer
func read_tag(tag byte, buf *bytes.Buffer)(n int, err error){
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