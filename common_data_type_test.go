package common_data_type

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNullData(t *testing.T) {
	buf := new(bytes.Buffer)
	value := NullData{}
	expect := []byte{0}
	got := value.Encode()
	var n int

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("wrong encode. got %b, expected %b", got, expect)
	}
	n, _ = value.WriteToBuffer(*buf)
	if n != 1 {
		t.Error("error write to buffer")
	}else{
		t.Logf("buffer: %v", buf)
	}
}
