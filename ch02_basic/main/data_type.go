package main

import "fmt"

func main() {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 // error
	fmt.Println(a, aPtr)
	fmt.Printf("%T, %T\n", a, aPtr)
}
