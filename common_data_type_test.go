package common_data_type

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeLength(t *testing.T) {
	values := []uint32{1, 0x7f, 0x80, 0x200, 0x3_03_00, 0xff_04_04_00}
	expected := [][]byte{{1}, {0x7f}, {0x81, 0x80}, {0x82, 2, 0}, {0x83, 3, 3, 0}, {0x84, 0xff, 4, 4, 0}}

	for i, value := range values {
		got := encode_length(value)
		if !reflect.DeepEqual(got, expected[i]) {
			t.Errorf("got %v, expect %v", got, expected[i])
			return
		}
		t.Logf("len %d, encode %v", value, got)
	}
}

func TestDecodeLength(t *testing.T) {
	values := [][]byte{{0}, {0x7f}, {0x80}, {0x81, 128}, {0x82, 2, 0}}
	expected := []uint32{0, 127, 0, 128, 512}
	for i, value := range values {
		buf := bytes.NewBuffer(value)
		got, _ := decode_length(buf)
		t.Log(got)
		if !reflect.DeepEqual(got, expected[i]) {
			t.Errorf("got %v, expect %v", got, expected[i])
			return
		}
	}
}

func TestNullData(t *testing.T) {
	buf := new(bytes.Buffer)
	value := new(NullData)
	expect := []byte{0}
	got := Encode(value)
	var n int

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("wrong encode. got %d, expected %d", got, expect)
	}
	n, _ = CDTtoBuffer(value, buf)
	if n != 1 {
		t.Error("error write to buffer")
	} else if buf.Len() != 1 {
		t.Errorf("expected length 1, got %d", buf.Len())
	} else {
		t.Logf("buffer: %v", buf.Bytes())
	}
}

func TestInteger(t *testing.T) {
	buf := new(bytes.Buffer)
	value := new(Integer)
	value.SetFromInt8(-1)
	expect := []byte{15, 0xff}

	got := Encode(value)
	var n int
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("wrong encode. got %d, expected %d", got, expect)
	}
	n, _ = CDTtoBuffer(value, buf)
	if n != 2 {
		t.Error("error write to buffer")
	} else if buf.Len() != 2 {
		t.Errorf("expected length 2, got %d", buf.Len())
	} else {
		t.Logf("buffer: %v", buf.Bytes())
	}
	value.Set(buf)
}

func TestEnum(t *testing.T) {
	buf := new(bytes.Buffer)
	value := new(Enum)
	value.SetFromByte(1)
	expect := []byte{22, 1}

	got := value.Encode()
	var n int
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("wrong encode. got %d, expected %d", got, expect)
	}
	n, _ = CDTtoBuffer(value, buf)
	if n != 2 {
		t.Error("error write to buffer")
	} else if buf.Len() != 2 {
		t.Errorf("expected length 2, got %d", buf.Len())
	} else {
		t.Logf("buffer: %v", buf.Bytes())
	}
	// value.Set([]byte{0})
}

// Todo: separate to other module
func TestUnit(t *testing.T) {
	buf := new(bytes.Buffer)
	value := new(UnitEnum)
	if err := value.SetFromByte(1); err != nil {
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
	} else if buf.Len() != 2 {
		t.Errorf("expected length 2, got %d", buf.Len())
	} else {
		t.Logf("buffer: %v", buf.Bytes())
	}
	// value.Set([]byte{0})
}

func TestOctetString(t *testing.T) {
	obj := new(OctetString)
	obj.SetFromString("привет")
	{
		expect := []byte{208, 191, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130}
		if !reflect.DeepEqual(expect, obj.Decode()) {
			t.Errorf("got %v expect %v", obj.Decode(), expect)
		}
	}
	{
		expect := []byte{9, 12, 208, 191, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130}
		got := obj.Encode()
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("got %v expect %v", obj.Decode(), expect)
		}
	}
}

func TestScalerUnit(t *testing.T) {
	value := new(ScalerUnit)
	a := value.Contents()
	t.Logf("%v", a)
	value.Scaler.SetFromInt8(3)
	value.Unit.SetFromByte(4)
	a1 := value.Encode()
	t.Logf("%v", a1)
	buf := bytes.NewBuffer(a1)
	value1 := new(ScalerUnit)
	value1.Set(buf)
	t.Log(value1)
}

func TestArray(t *testing.T) {
	buf := bytes.NewBuffer([]byte{1, 2, 15, 10, 15, 11})
	value := new(IntegerArray)
	err := value.Set(buf)
	t.Log(value.Encode(), err)
}

func TestLong(t *testing.T) {
	value := new(Long)
	buf := bytes.NewBuffer([]byte{16, 1, 10})
	Set(value, buf)
	var expect int16 = 266
	got := value.ToInt16()
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("got %v expect %v", got, expect)
	}
	a := Encode(value)
	t.Log(Encode(value), a)
}
