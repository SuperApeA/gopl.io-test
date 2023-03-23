package weightconv

import "fmt"

const Weight = "weight"

type Pound float64
type Kilogram float64

func (p Pound) String() string    { return fmt.Sprintf("%g lb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }

//!-
