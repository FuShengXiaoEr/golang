package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(defFIleName string,srcFileName string)(written int64,err error) {
	srcFile,err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file error = %v\n",err)
		return 0,err
	}
	defer srcFile.Close()
	read := bufio.NewReader(srcFile)

	desFile,err := os.OpenFile(defFIleName,os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		 fmt.Printf("open file error = %v\n",err)
	}
	defer desFile.Close()
	write := bufio.NewWriter(desFile)
	return io.Copy(write,read)
}

func main(){
	CopyFile("abc_copy.txt","abc.txt")
}