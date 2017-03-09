package main

import (
	"fmt"
	"os"
)

func main() {
	for i, line := range os.Args[:] {
		fmt.Println(i, line)
	}
}
