package main

import (
	"fmt"
	"sort"
)

func main(){
	var intSlice = []int{0,-1,10,7,90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}