package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("This is the %v parm: %v\n", i, os.Args[i])
	}
}
