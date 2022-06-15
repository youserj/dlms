package common_data_type

import (
	"bytes"
)

type NullData struct{}

func (NullData) TAG() byte {
	return 0
}

func (NullData) Contents() []byte {
	return make([]byte, 0)
}

func (c NullData) Encode() (ret []byte) {
	ret = []byte{c.TAG()}
	return
}

func (c *NullData) SetFromBuffer(buf bytes.Buffer) (n int, err error) {
	return read_tag(c.TAG(), buf)
}