package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

func main() {
	//a, _ := returnMultiValues()
	//fmt.Println(a)

	//tsSF := timeSpent(slowFunc)
	//fmt.Println(tsSF(10))

	//fmt.Println(sum(1, 2, 3)) // 6

	//deferDemo()

	initEmployee()
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}

func deferDemo() {
	defer func() {
		fmt.Println("deferDemo defer")
	}()
	fmt.Println("deferDemo")
	panic("Fatal error") // defer函数仍会执行
}

type Employee struct {
	Id   string
	Name string
	Age  int
}

func initEmployee() {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Id: "1", Name: "Mike", Age: 21}
	e2 := new(Employee) // 注意这里返回的引用/指针，相当于 e := &Employee{}
	e2.Id = "2" // 与其它语言的差异，通过实例的指针访问不需要使用 ->
	e2.Name = "Rose"
	e2.Age = 18
	fmt.Println(e)
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Printf("%T, %T, %T\n", e, e1, e2)
}

// 第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) String1() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %s - Name: %s - Age: %s", e.Id, e.Name, e.Age)
}

// 通常情况下为了避免内存拷贝使用第二种定义方式
func (e *Employee) String2() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %s - Name: %s - Age: %s", e.Id, e.Name, e.Age)
}