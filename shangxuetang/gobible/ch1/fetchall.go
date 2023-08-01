package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for n,url := range os.Args[1:] {
		fmt.Println(n,url)
		go fetch2(url, ch,n)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err :=io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}

func fetch2(url string,ch chan<- string,n int) {
	start := time.Now()
	rep,err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	filename := "fetchall_test"+strconv.Itoa(n)+".txt"
	fmt.Printf("filename: %s\n",filename)
	out, err := os.Create(filename)
	write := bufio.NewWriter(out)

	nbytes, err := io.Copy(write, rep.Body)
	rep.Body.Close()

	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	write.Flush()

}