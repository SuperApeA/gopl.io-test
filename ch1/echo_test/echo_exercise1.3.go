package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	method1Start := time.Now()
	str, sep := "", ""
	for _, arg := range os.Args[1:] {
		str += sep + arg
		sep = " "
	}
	fmt.Printf("First method use time: %v\n", time.Now().Sub(method1Start))
	fmt.Println(str)
	method2Start := time.Now()
	str = strings.Join(os.Args[1:], " ")
	fmt.Printf("Second method use time: %v\n", time.Now().Sub(method2Start))
	fmt.Println(str)
}
