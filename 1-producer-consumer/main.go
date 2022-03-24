//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync" //<!!> Adding sync
	"time"
)

//<!!> Change the args to send in the channel to send tweets on
// and also the waitgroup (object?)
func producer(stream Stream, events chan *Tweet, wg *sync.WaitGroup) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			wg.Done() //<!!> If EOF, producer done, set wg to done. And Break.
			break
		} else {
			//<!!> Else broadcast tweet to the channel
			events <- tweet
		}
	}
	//<!!> Close the channel
	close(events)
}

func consumer(events chan *Tweet, wg *sync.WaitGroup) {
	//<!!> Read from the channel till it is open.
	for t := range events {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	//<!!> Once closed. Mark the wg as done.
	wg.Done()
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	events := make(chan *Tweet)

	//<!!> Maintain a wg for 2 routines
	var wg sync.WaitGroup
	wg.Add(2)

	// Producer
	go producer(stream, events, &wg)
	// Consumer
	go consumer(events, &wg)

	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
