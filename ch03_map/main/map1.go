package main

import "fmt"

func main() {
	//mapOp()
	//mapFunc()
	map2Set()
}

func mapOp() {
	m1 := map[string]int{}
	fmt.Println(m1["haha"]) // default 0
	//m1["haha"] = 2
	if v, ok := m1["haha"]; ok {
		fmt.Println("exist haha:", v)
	} else {
		fmt.Println("don't exist")
	}
}

func mapFunc() {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	fmt.Println(m[1](2), m[2](2), m[3](2)) // 2 4 8
}

func map2Set() {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		fmt.Printf("%d exist\n", n)
	} else {
		fmt.Printf("%d not exist\n", n)
	}
	mySet[3] = true
	fmt.Println(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		fmt.Printf("%d exist\n", n)
	} else {
		fmt.Printf("%d not exist\n", n)
	}
}
