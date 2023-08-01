package main

import "fmt"

func printyramid(level int){
	for i := 1;i <= level ;i++ {
		for k := 1 ;k <= level - i;k++ {
			fmt.Print(" ")
		}

		for m := 1; m <= i*2 -1;m++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	
}

func main() {
	printyramid(9)
}