package goplay

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sendAlert() {
	fmt.Println("Sending alert...")
	done := make(chan struct{})
	go func() {
		defer close(done)
		time.Sleep(1 * time.Second)
		fmt.Println("Alert sent!")
		done <- struct{}{}

	}()
	<-done
}

func scoping() {
	printI := func(i int) {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		go printI(i)
	}
}

func fetch(s string) string { return s }

func dataFetcher(s string, ch chan<- string) {
	ch <- fetch(s)
}

func leaker(ctx context.Context) {
	expectedInputs := []string{"https://www.google.com", "https://www.facebook.com", "https://www.twitter.com"}
	ch := make(chan string, len(expectedInputs))
	go dataFetcher("https://www.google.com", ch)

	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-ctx.Done():
			fmt.Println("Context deadline: Exited!")
		}
	}
}

func normalMap() {
	key := "key"
	var wg sync.WaitGroup
	commonMap := map[string]int{key: 1}
	mapSetter := func(wg *sync.WaitGroup, i int) {
		defer wg.Done()
		commonMap[key] = i
	}
	// mapReader := func() {
	// 	fmt.Println(commonMap[key])
	// }

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mapSetter(&wg, i)
		// go mapReader()
	}
	wg.Wait()
}
