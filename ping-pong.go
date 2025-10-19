package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}

}

func printar(c chan string) {
	msg := <-c
	fmt.Println(msg)
	time.Sleep(time.Second * 1)
}

func main() {
	c := make(chan string)

	go ping(c)
	go printar(c)
	go pong(c)

	var saida string
	fmt.Scanln(&saida)

}
