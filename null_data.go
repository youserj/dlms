package common_data_type

import "bytes"

type NullData struct{}

func (*NullData) TAG() byte {
	return 0
}

func (*NullData) ContentsLen() uint32 {
	return 0
}

func (*NullData) Length() []byte {
	return []byte{}
}

func (NullData) Contents() []byte {
	return make([]byte, 0)
}

func (c *NullData) Encode() (ret []byte) {
	ret = []byte{c.TAG()}
	return
}

func (c *NullData) Set(buf *bytes.Buffer) error{ 
	return read_tag(c, buf)
}