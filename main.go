package main

import (
	"fmt"
	"sync"
	"time"
)

func send(c chan<- int, i int, wg *sync.WaitGroup) {
	fmt.Println("Send a job", i)
	c <- i
	time.Sleep(time.Millisecond * 1000)

	wg.Done()
}

func check(c <-chan int) {
	for {
		fmt.Println("receive a job", <-c)
	}
}

func receive(c <-chan int) {

	for {
		fmt.Println(<-c)
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go send(c, i, &wg)
	}
	go func() {
		check(c)
		wg.Done()
		close(c)
	}()

	wg.Wait()
}
