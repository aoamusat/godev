package main

type Celcius float64
type Fahrenheit float64

const AbsoluteZeroC Celcius = -273.15
const FreezingC Celcius = 0
const BoilingC Celcius = 100

func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }
