package tempconv

//CtoF converts Celsius to Fahrenheit
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

// FtoC converts Fahrenheit to Celsius
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f -32) * 5 / 9)
}

func KtoC(k Kelvin) Celsius {
	return Celsius(k + AbsoluteZeroC)
}

func CtoK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}