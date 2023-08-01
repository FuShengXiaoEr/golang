package main

import "fmt"
func main (){
	xMap := map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}
	yMap := map[string]int{
		"a":1,
		"b":2,
		"c":4,
	}
	fmt.Printf("是否相等%t",equal(xMap,yMap))
}

func equal(x ,y map[string]int)bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv,ok := y[k]; !ok || yv != xv{
			return false
		}
	}

	return true
}