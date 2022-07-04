package common_data_type

import "bytes"

type OctetString struct {
	contents []byte
}

func (*OctetString) TAG() byte {
	return 9
}

func (t *OctetString) ContentsLen() uint32 {
	return uint32(len(t.contents))
}

// TODO: may be use contents copy
func (t *OctetString) Contents() []byte {
	return t.contents
}

func (t *OctetString) Encode() []byte {
	ret := []byte{t.TAG()}
	ret = append(ret, encode_length(t.ContentsLen())...)
	ret = append(ret, t.Contents()...)
	return ret
}

func (t *OctetString) Set(buf bytes.Buffer) error {
	// not implement now
	return nil
}

func (t *OctetString) SetFromString(value string) {
	t.contents = []byte(value)
}

func (t OctetString) Decode() []byte {
	return t.contents
}
