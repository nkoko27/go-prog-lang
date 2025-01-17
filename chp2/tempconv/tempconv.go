package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	Boiling Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%gC", c)}
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f)}
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k)}

// Exercise 2.1
// add Kelvin