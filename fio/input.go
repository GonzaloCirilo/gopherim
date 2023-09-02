package fio

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var Scanner = bufio.NewScanner(os.Stdin)

type Integer interface {
	~int | ~int32 | ~int64
}

func toInt[I Integer](buf []byte) (n I) {
	for _, v := range buf {
		n = n*10 + I(v-'0')
	}
	return
}

func toString(buf []byte) string {
	return string(buf)
}

func toFloat64(buf []byte, precision int) (f float64) {
	s := string(buf)
	f, err := strconv.ParseFloat(s, precision)
	if err != nil {
		panic("Error parsing float64")
	}
	return
}

// Scanf Scans variables given a format. Returns false if EOF reached
func Scanf(format string, a ...any) bool {
	
	tokens := strings.Split(format, "%")
	if len(tokens) > len(a) {
		panic("Not enough variables specified")
	} else if len(tokens) < len(a) {
		panic("")
	}
	for i, t := range tokens {
		scanSuccess := Scanner.Scan()
		if scanSuccess == false && Scanner.Err() == nil {
			return false
		} else if scanSuccess == false && Scanner.Err() != nil{
			continue
		}
		b := Scanner.Bytes()
		switch t {
		case "d":
			a[i] = toInt[int](b)
		case "f":
			a[i] = toFloat64(b, 8)
		case "s":
			a[i] = toString(b)
		default:
			a[i] = b
		}
	}
	return true
}

// Reads an integer from standard input. Specify type of interger.
// Supported types are int, int32 & int64
func ReadInteger[I Integer]() (I, bool) {
	scanSuccess := Scanner.Scan()
	isEof := !scanSuccess && Scanner.Err() == nil
	b := Scanner.Bytes()
	return toInt[I](b), !isEof
}

func ReadString() (string, bool) {
	scanSuccess := Scanner.Scan()
	isEof := !scanSuccess && Scanner.Err() == nil
	b := Scanner.Bytes()
	return toString(b), !isEof
}

func ReadFloat64(precision int) (float64, bool) {
	scanSuccess := Scanner.Scan()
	isEof := !scanSuccess && Scanner.Err() == nil
	b := Scanner.Bytes()
	return toFloat64(b, precision), !isEof
}
