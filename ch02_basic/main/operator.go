package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	fmt.Println(a == b)
	//fmt.Println(a == c) // error
	fmt.Println(a == d)
}
