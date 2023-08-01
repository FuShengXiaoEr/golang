package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main(){
	// fortest()
	// stringJoinTest()
	time.
}

func fortest(){
	s := ""
	for  i := 1 ; i<len(os.Args);i++ {
		fmt.Printf("i=%d,value=%v",i,os.Args[i])
		fmt.Println()
	}
	fmt.Println(s)
}

func stringJoinTest(){
	fmt.Println(strings.Join(os.Args[0:]," "))
}

func echoString(){
	s,sep := "",""
	for  _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

