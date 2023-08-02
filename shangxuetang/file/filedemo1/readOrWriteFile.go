package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fileName := "abc.txt"
	file,err := os.OpenFile(fileName,os.O_CREATE | os.O_WRONLY,0666)
	if err != nil {
		fmt.Printf("open file error = %v\n",err)
		return
	}

	defer file.Close()
	str :="hello god\n"
	write :=  bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString(str)
		
	}
	write.Flush()

}