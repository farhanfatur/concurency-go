package main

import (
	"fmt"
	"sync"
	"time"
)

func send(c chan<- int, i int, wg *sync.WaitGroup) {
	fmt.Println("send a job", i)
	c <- i
	time.Sleep(time.Millisecond * 1000)

	wg.Done()
}

func receive(c <-chan int, cReceive chan int) {
	for {
		var num = <-c
		select {
		case <-c:
			fmt.Println("Receive a job", <-c)
		case cReceive <- <-c:
			cReceive <- num
			fmt.Println(<-cReceive)
		}
	}

}

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	c := make(chan int)
	cReceive := make(chan int)
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go send(c, i, &wg)
	}
	go func() {
		mtx.Lock()
		receive(c, cReceive)
		mtx.Unlock()
	}()
	fmt.Println(cReceive)
	// go func() {
	// 	for {
	// 		fmt.Println("parsingReceive:", cReceive)
	// 	}
	// }()
	wg.Wait()
}
