package main

import (
	"fmt"
	"time"
)

func service() string {
	time.Sleep(time.Microsecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("Task is done.")
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1) // 更高效，不阻塞
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret  // retCh buf 为 1 不阻塞
		fmt.Println("service existed")
	}()
	return retCh
}

func main() {
	//fmt.Println(service())
	//otherTask()

	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second + 1)
}
