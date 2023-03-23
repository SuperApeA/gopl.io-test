// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/cf_test/lengthconv"
	"gopl.io/ch2/cf_test/weightconv"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case lengthconv.Length:
			for _, arg := range args[1:] {
				l, err := strconv.ParseFloat(arg, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "fm: %v\n", err)
					os.Exit(1)
				}
				f := lengthconv.Feet(l)
				m := lengthconv.Meter(l)
				fmt.Printf("%s = %s, %s = %s\n",
					f, lengthconv.FToM(f), m, lengthconv.MToF(m))
			}
		case weightconv.Weight:
			for _, arg := range args[1:] {
				w, err := strconv.ParseFloat(arg, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "pk: %v\n", err)
					os.Exit(1)
				}
				p := weightconv.Pound(w)
				k := weightconv.Kilogram(w)
				fmt.Printf("%s = %s, %s = %s\n",
					p, weightconv.PToK(p), k, weightconv.KToP(k))
			}
		default:
			fmt.Println("type is invalid")
			return
		}
	}
}

//!-
