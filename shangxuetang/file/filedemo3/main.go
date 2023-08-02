package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type item struct {
	chCount int
	numCount int
	spaceCount int
	otherCount int
}

func main(){
	readFile ,err := os.Open("abc.txt")
	if err != nil {
		fmt.Printf("open file error =%v\n",err)
		return
	}
	defer readFile.Close()

	var count item

	reader := bufio.NewReader(readFile)

	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
			
		}
		fmt.Printf("str type %T\n",str)
		
		for _,v := range str {
			if v >= 'a' && v <= 'z' {
				count.chCount++
			}else if v >= 'A' && v <= 'Z' {
				count.chCount++ 
			}else if v == ' ' || v <= '\t' {
				count.spaceCount++ 
			}else if v >= '0' && v <= '9' {
				count.numCount++ 
			}else {
				count.otherCount++
			}
		}
	}

	fmt.Printf("finally,count chCount:%d,spaceCount:%d,numCount:%d,otherCount:%d",count.chCount,
	count.spaceCount,count.numCount,count.otherCount)
}