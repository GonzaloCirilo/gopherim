package fio

import (
	"bufio"
	"fmt"
	"os"
)

var writer = bufio.NewWriter(os.Stdout)

func Printf(format string, a ...any) {
	fmt.Fprintf(writer, format, a...)
}

func Println[K any](x K) {
	fmt.Fprintln(writer, x)
}

func Flush() {
	writer.Flush()
}