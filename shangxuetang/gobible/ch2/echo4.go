package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "换行，默认是不换行")
var sep = flag.String("s"," ","分隔符")

func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}