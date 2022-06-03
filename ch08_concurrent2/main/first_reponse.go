package main

import (
	"fmt"
	"runtime"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Microsecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner) // 改成 buffer channel，不用阻塞等待
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner) // 改成 buffer channel，不用阻塞等待
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for i:=0 ; i<numOfRunner; i++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func main() {
	fmt.Println("Before:", runtime.NumGoroutine())
	//fmt.Println(FirstResponse())
	fmt.Println(AllResponse())
	fmt.Println("After:", runtime.NumGoroutine())
}
