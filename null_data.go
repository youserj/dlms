package common_data_type

type NullData struct{}

func (*NullData) TAG() byte {
	return 0
}

func (*NullData) ContentsLen()int{
	return 0
}

func (NullData) Contents() []byte {
	return make([]byte, 0)
}

func (c *NullData) Encode() (ret []byte) {
	ret = []byte{c.TAG()}
	return
}

func (*NullData) Set(_ []byte)(err error){
	return
}