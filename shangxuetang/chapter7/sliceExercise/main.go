package main

import "fmt"

func main(){
	var a []int = []int{1,2,3,4,5}
	var slice = make([]int,1)
	fmt.Println(slice)
	copy(slice,a)
	fmt.Println(slice)

	
	fnbslice := fbn(10)
	fmt.Println(fnbslice)
}

/**
可以接受一个n int
能狗接受斐波那契数列放到数列中

**/
func fbn(n int) ([]uint64){
	fbnslice := make([]uint64,n)
	fbnslice[0] = 1
	fbnslice[1] = 1
	for i := 2 ;i < n;i++ {
		fbnslice[i] = fbnslice[i-1] + fbnslice[i-2]
	}
	return fbnslice
}