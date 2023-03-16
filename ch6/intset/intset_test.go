// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestIntSet(t *testing.T) {
	a := assert.New(t)
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	a.Equal(4, x.Len())
	x.Remove(1)
	a.Equal(3, x.Len())
	y := x.Copy()
	a.Equal(true, y.Has(144))
	a.Equal(true, y.Has(9))
	a.Equal(true, y.Has(42))
	a.Equal(3, y.Len())
	x.Clear()
	a.Equal(0, x.Len())
	a.Equal(3, y.Len())
	x.AddAll(10, 11, 12, 13, 14)
	a.Equal(5, x.Len())
	a.Equal(3, y.Len())
}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	y.Add(1)
	y.Add(145)
	y.Add(10)
	y.Add(43)

	x.IntersectWith(&y)
	fmt.Println(x.String())
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	y.Add(1)
	y.Add(145)
	y.Add(10)
	y.Add(43)

	x.DifferenceWith(&y)
	fmt.Println(x.String())
}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	y.Add(1)
	y.Add(145)
	y.Add(10)
	y.Add(43)

	x.SymmetricDifference(&y)
	fmt.Println(x.String())
}

func TestElems(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(x.Elems())
}
