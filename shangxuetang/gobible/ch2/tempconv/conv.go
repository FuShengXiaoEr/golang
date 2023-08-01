package tempconv


// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c tempconv.Fahrenheit) { return tempconv.Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f tempconv.Fahrenheit) tempconv.Celsius { return tempconv.Celsius((f - 32) * 5 / 9) }