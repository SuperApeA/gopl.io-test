package lengthconv

import "fmt"

const Length = "length"

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }

//!-
