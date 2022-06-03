package main

import (
	"fmt"
	"sync"
)

func fun() {
	//// 多渠道的选择
	//select {
	//case ret := <-retCh1:
	//	fmt.Println("xxx")
	//case ret := <-retCh2:
	//	fmt.Println("xxx")
	//default:
	//	fmt.Println("xxx")
	//}
	//
	//// 超时控制
	//select {
	//case ret := <-retCh:
	//	fmt.Println("xxx")
	//case <-time.After(time.Second * 1)
	//	fmt.Println("time out")
	//}
}

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		// ch <- 11 // error panic
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		//for i := 0; i < 10; i++ {
		//	data := <-ch
		//	fmt.Println(data)
		//}
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
