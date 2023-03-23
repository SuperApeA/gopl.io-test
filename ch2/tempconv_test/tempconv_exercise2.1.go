package main

import "fmt"

func main() {
	c := Celsius(100)
	fmt.Printf("100°C is equal %s\n", CToF(c).String())
	fmt.Printf("100°C is equal %s\n", CToK(c).String())

	f := Fahrenheit(100)
	fmt.Printf("100°F is equal %s\n", FToC(f).String())
	fmt.Printf("100°F is equal %s\n", FToK(f).String())

	k := Kelvin(100)
	fmt.Printf("100 K is equal %s\n", KToC(k).String())
	fmt.Printf("100 K is equal %s\n", KToF(k).String())
}
