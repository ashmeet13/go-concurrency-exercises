package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			<-interrupt
			fmt.Println("Here")
		}
	}()

	time.Sleep(10 * time.Second)
}
