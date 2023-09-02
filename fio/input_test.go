package fio

import (
	"strconv"
	"testing"
)

func TestToInt(t *testing.T) {
	bs := []byte(strconv.Itoa(31415926))
	a := toInt[int](bs)
	if a != 31415926 {
		t.Fatalf("Expected 31415926 actual value %d", a)
	}
}

func TestToInt32(t *testing.T) {
	bs := []byte(strconv.Itoa(45))
	a := toInt[int32](bs)
	if a != 45 {
		t.Fatalf("Expected 45 actual value %d", a)
	}
}

func TestToInt64Fails(t *testing.T) {
	bs := []byte(strconv.Itoa(40))
	a := toInt[int64](bs)
	if a == 45 {
		t.Fatalf("Expected different values")
	}
}

func TestToString(t *testing.T) {
	expected := "This is a string"
	bs := []byte(expected)
	s := toString(bs)
	if s != expected {
		t.Fatalf("Expected: %s actual value: %s", expected, s)
	}
}

func TestToFloat64(t *testing.T) {
	input := "20005.37829473"
	expected := 20005.37829473
	bs := []byte(input)
	f := toFloat64(bs, 8)
	if f != expected {
		t.Fatalf("Expected: %f actual value: %f", expected, f)
	}
}

func TestToFloat64With6Precision(t *testing.T) {
	input := "20005.378294"
	expected := 20005.37829400
	bs := []byte(input)
	f := toFloat64(bs, 6)
	if f != expected {
		t.Fatalf("Expected: %f actual value: %f", expected, f)
	}
}
