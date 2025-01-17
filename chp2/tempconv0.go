// Package tempconv performs Ceslsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Fahrenheit float64
type Celsius float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5/9)
}

func (c Celsius) String() string { return fmt.Sprintf("%gC", c)}