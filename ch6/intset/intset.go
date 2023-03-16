// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
	"math/bits"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-n	egative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// add some bit
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(listInt ...int) {
	for _, x := range listInt {
		s.Add(x)
	}
}

// Remove remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// remove some bit
	s.words[word] &= ^(uint64(1) << bit)
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith 交集：元素在A集合B集合均出现
func (s *IntSet) IntersectWith(t *IntSet) {
	if s.Len() > t.Len() {
		s.words = s.words[0:t.Len()]
	}
	for i, twords := range t.words {
		s.words[i] &= twords
	}
}

// DifferenceWith 差集：元素出现在A集合，未出现在B集合
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, twords := range t.words {
		if i > s.Len() {
			return
		}
		s.words[i] &= ^twords
	}
}

// SymmetricDifference 并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A
func (s *IntSet) SymmetricDifference(t *IntSet) {
	intersectSet := s.Copy()
	intersectSet.IntersectWith(t)
	s.UnionWith(t)
	s.DifferenceWith(intersectSet)
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len return the number of elements
func (s *IntSet) Len() int {
	size := 0
	for _, x := range s.words {
		size += bits.OnesCount64(x)
	}
	return size
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	if len(s.words) == 0 {
		return
	}
	s.words = make([]uint64, 0)
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	newIntSet := new(IntSet)
	newIntSet.words = make([]uint64, len(s.words))
	copy(newIntSet.words, s.words)
	return newIntSet
}

// Elems 返回集合中的所有元素
func (s *IntSet) Elems() []uint64 {
	list := make([]uint64, 0, s.Len())
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				list = append(list, uint64(64*i+j))
			}
		}
	}
	return list
}
