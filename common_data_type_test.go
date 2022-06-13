package common_data_type

import (
	"testing"
	"reflect"
	// "bytes"
)

func TestNullData(t *testing.T) {
	value := NullData{}
	expect := []byte{0}
	got := value.Encode()
	if !reflect.DeepEqual(expect, got){
		t.Errorf("wrong encode. got %b, expected %b", got, expect)
	}
}