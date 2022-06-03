package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC() // 清除 sync.Pool 中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	v2, _ := pool.Get().(int)
	fmt.Println(v2)
}
