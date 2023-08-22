package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func notify(services ...string) {
	var wg sync.WaitGroup
	wg.Add(len(services))

	for a, service := range services {
		go func(a int, s string) {
			fmt.Printf("Starting service-%d: %s\n", a, s)
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			fmt.Printf("Completed service-%d: %s\n", a, s)
			wg.Done()
		}(a, service)
	}
	wg.Wait()
	fmt.Println("All services started!")
}

func notifyChannels(services ...string) {

	resp := make(chan string)
	count := 0

	for _, service := range services {
		count++
		go func(s string) {
			fmt.Printf("Start notifying service:%s\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			resp <- "Completed notifying service:" + s
		}(service)
	}

	for i := 0; i < count; i++ {
		fmt.Println(<-resp)
	}
}

func main() {

	notify("Service-1", "Service-2", "Service-3")
	notifyChannels("Service-1", "Service-2", "Service-3")
}
