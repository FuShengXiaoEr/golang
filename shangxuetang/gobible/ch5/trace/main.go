package main

import (
	"log"
	"time"
)

func main(){
	defer trace("bigSlowOperation")()
	time.Sleep(time.Second*5)

}

func trace(msg string)func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)",msg, time.Since(start))
	}
}