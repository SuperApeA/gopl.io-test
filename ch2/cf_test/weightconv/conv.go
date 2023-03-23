package weightconv

// PToK converts a Pound temperature to Kilogram.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KToP converts a Kilogram temperature to Pound.
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }
