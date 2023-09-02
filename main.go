package main

import (
	"fmt"
	"io"
	"time"

	"github.com/GonzaloCirilo/gopherim/fio"
)

func fastInput() {
	for {
		_, canScan := fio.ReadInteger[int]()
		if canScan == false {
			break
		}
		
	}

}

func regularInput() {
	for {
		var x int
		_, err := fmt.Scanf("%d", &x)
		if err == io.EOF {
			break
		}
	}
}

func main() {
	start := time.Now()
	regularInput()
	elapsed := time.Since(start)
	fmt.Printf("Exec time %s\n", elapsed)
}
