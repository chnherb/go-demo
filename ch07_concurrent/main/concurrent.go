package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//mutex()
	//waitGroup()
	mutexAndWaitGroup()
}

func mutexAndWaitGroup() {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func waitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
		}()
	}
	wg.Wait()
}

func mutex() {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(counter)
}
