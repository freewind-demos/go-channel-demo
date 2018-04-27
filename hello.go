package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(100)

	channel := make(chan string)
	go consumeData(channel)
	go produceData("A", channel)
	go produceData("B", channel)

	wg.Wait()
	fmt.Println("Exit")
}

func consumeData(channel chan string) {
	for {
		data := <-channel
		fmt.Println(data)
		wg.Done()
	}
}

func produceData(label string, channel chan string) {
	for {
		data := fmt.Sprintf("%s: %d", label, rand.Int())
		fmt.Println(data)
		channel <- data
	}
}
