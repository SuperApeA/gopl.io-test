package lengthconv

// FToM converts a Feet temperature to Meter.
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

// MToF converts a Meter temperature to Feet.
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

//!-
