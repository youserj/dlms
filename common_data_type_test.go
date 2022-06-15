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
	n, _ = CDTtoBuffer(value, buf)
	if n != 1 {
		t.Error("error write to buffer")
	}else if buf.Len() != 1{
		t.Errorf("expected length 1, got %d", buf.Len())
	}else{
		t.Logf("buffer: %v", buf.Bytes())
	}
}

func TestInteger(t *testing.T) {
	buf := new(bytes.Buffer)
	value := Integer{1}
	expect := []byte{15, 1}
	got := value.Encode()
	var n int

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("wrong encode. got %b, expected %b", got, expect)
	}
	n, _ = CDTtoBuffer(value, buf)
	if n != 2 {
		t.Error("error write to buffer")
	}else if buf.Len() != 2{
		t.Errorf("expected length 2, got %d", buf.Len())
	}else{
		t.Logf("buffer: %v", buf.Bytes())
	}
}
