package main

type Celsius float64
type Fahrenheit float64

const(
	AbsoultZeroC Celsius =  -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func main(){
	CToF()

}

func CToF(c Celsius) Fahrenheit{
	return Fahrenheit(c*9/5 + 32)
}

func FToC(F Fahrenheit) Celsius{
	return Celsius((f-32)*5/9)
}