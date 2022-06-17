package common_data_type

type Integer int8

func (*Integer) TAG()byte{
	return 15
}

func (*Integer) ContentsLen()int{
	return 1
}

func (c Integer) Contents()(ret []byte){
	ret = []byte{byte(c)}
	return
}

func (c *Integer) Encode()(ret []byte){
	ret = []byte{c.TAG(), byte(*c)}
	return
}

func (c *Integer) Set(value []byte)(err error){
	*c = Integer(value[0])
	return
}