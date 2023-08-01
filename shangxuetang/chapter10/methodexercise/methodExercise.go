package main

import "fmt"


type Circle struct {
	radius float64
	
}

func (c Circle)area() float64 {
	return 3.14 * c.radius*c.radius
}

func (cir Circle)String() string{
	return fmt.Sprintf("radio~%f",cir.radius)
}

func main(){
	var c Circle
	c.radius = 4.0
	res := c.area()
	fmt.Println(res)

	var cir = &Circle{
		radius: 7.6,
	}

	fmt.Println(*cir)
}

