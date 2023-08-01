package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
func main() {
	main3()
}


func main1() {
	for _,url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil  {
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}
}

func main2(){
	// 获取传入参数的url
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		out, err := os.Create("out.txt")
		wt := bufio.NewWriter(out)

		b, err := io.Copy(wt,resp.Body)
		fmt.Println("write:",b)
		if err != nil  {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
		}
		wt.Flush()
	}
}


func main3(){
	// 获取传入参数的url
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") == false { 
			url = "http://" + url
		}
		resp, err := http.Get(url)
		fmt.Println("http code",resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		out, err := os.Create("out.txt")
		wt := bufio.NewWriter(out)

		b, err := io.Copy(wt,resp.Body)
		fmt.Println("write:",b)
		if err != nil  {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
		}
		wt.Flush()
	}
}