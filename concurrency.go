package goplay

import (
	"fmt"
	"sync"
	"time"
)

func leaker(timeout time.Duration) {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go dataFetcher("https://www.google.com", ch, &wg)

	select {
	case data := <-ch:
		fmt.Println(data)
	case <-time.After(timeout):
		fmt.Println("Context deadline: Exited!")
	}

	wg.Wait()
}

func dataFetcher(url string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	ch <- url
}
