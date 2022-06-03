package main

import (
	"fmt"
	"os"
)

func main() {
	// go run hello.go
	// go build hello.go
	// ./hello

	// go run hello.go haha
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("args:", os.Args[1])
	}
	fmt.Println("hello world")

	//os.Exit(-1) // exit status 255
	os.Exit(-1) // 0 success

}
