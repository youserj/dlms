package common_data_type

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNullData(t *testing.T) {
	buf := new(bytes.Buffer)
	value := new(NullData)
	expect := []byte{0}
	got := value.Encode()
	var n int

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("wrong encode. got %d, expected %d", got, expect)
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
	value := new(Integer)
	*value = 1
	expect := []byte{15, 1}
	
	// value.Set([]byte{2,4,5})
	got := value.Encode()
	var n int
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("wrong encode. got %d, expected %d", got, expect)
	}
	n, _ = CDTtoBuffer(value, buf)
	if n != 2 {
		t.Error("error write to buffer")
	}else if buf.Len() != 2{
		t.Errorf("expected length 2, got %d", buf.Len())
	}else{
		t.Logf("buffer: %v", buf.Bytes())
	}
	*value = 0
	BufferToCDT(value, buf)
	if buf.Len() != 0{
		t.Error("Buffer not empty")
	}
}

func TestEnum(t *testing.T) {
	buf := new(bytes.Buffer)
	value := new(Enum)
	*value = 1
	expect := []byte{22, 1}
	
	// value.Set([]byte{2,4,5})
	got := value.Encode()
	var n int
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("wrong encode. got %d, expected %d", got, expect)
	}
	n, _ = CDTtoBuffer(value, buf)
	if n != 2 {
		t.Error("error write to buffer")
	}else if buf.Len() != 2{
		t.Errorf("expected length 2, got %d", buf.Len())
	}else{
		t.Logf("buffer: %v", buf.Bytes())
	}
	*value = 0
	BufferToCDT(value, buf)
	if buf.Len() != 0{
		t.Error("Buffer not empty")
	}
}

// Todo: separate to other module
func TestUnit(t *testing.T) {
	buf := new(bytes.Buffer)
	value := new(Unit)
	if err := value.Set([]byte{1}); err != nil{
		t.Errorf("%s", err)
	}
	expect := []byte{22, 1}
	
	// value.Set([]byte{2,4,5})
	got := value.Encode()
	var n int
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("wrong encode. got %d, expected %d", got, expect)
	}
	n, _ = CDTtoBuffer(value, buf)
	if n != 2 {
		t.Error("error write to buffer")
	}else if buf.Len() != 2{
		t.Errorf("expected length 2, got %d", buf.Len())
	}else{
		t.Logf("buffer: %v", buf.Bytes())
	}
	value.Set([]byte{0})
	BufferToCDT(value, buf)
	if buf.Len() != 0{
		t.Error("Buffer not empty")
	}
}