package main

import (
	"fmt"
	"time"
)

func isCancelled(cancelCh chan struct{}) bool {
	select {
	case <-cancelCh:
		return true
	default:
		return false
	}
}

func cancel_1(cancelCh chan struct{}) {
	cancelCh <- struct{}{}
}

func cancel_2(cancelCh chan struct{}) {
	close(cancelCh)
}

func taskCancel() {
	cancelCh := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelCh)
	}
	//cancel_1(cancelCh)
	cancel_2(cancelCh)
	time.Sleep(time.Second * 1)
}

func main() {
	taskCancel()
}
