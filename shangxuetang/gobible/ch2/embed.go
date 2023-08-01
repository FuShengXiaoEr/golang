package main

import "fmt"

type Point struct {
	X,Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main(){
	var w Wheel
	w = Wheel{Circle{Point{8,8},5},20}
	w.X = 9
	w.Y = 9
	w.Radius = 6
	w.Spokes = 21
	fmt.Printf("%#v\n",w)
	
}