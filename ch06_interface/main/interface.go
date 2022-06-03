package main

import (
	"encoding/json"
	"fmt"
)

type Code string

type Programmer interface {
	WriteHelloWord() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWord() Code {
	return "fmt.Println(\"Hello world\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWord() Code {
	return "System.out.Println(\"Hello world\")"
}

func writeHelloWorld(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWord())
}

func main() {
	//goProg := new(GoProgrammer) // 指针类型，或为&GoProgrammer{}，但不能为 GoProgrammer{}
	//javaProg := new (JavaProgrammer)
	//writeHelloWorld(goProg)
	//writeHelloWorld(javaProg)

	//DoSomething(10)
	//DoSomething("10")

	fmt.Println(StructTran())
}

func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("int", i)
	//	return
	//}
	//if s, ok := p.(string); ok {
	//	fmt.Println("string", s)
	//	return
	//}
	//fmt.Println("Unkown Type")
	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unkown Type")
	}
}

type Info struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func StructTran() *Info {
	value := getInfo()
	//var info *Info
	if info, ok := value.(*Info); info != nil && ok {
		//fmt.Println(info.Name, info.ID)
		return info
	} else {
		fmt.Println("err")
	}
	return nil
}

func getInfo() interface{} {
	jsonStr := "{\"id\":1, \"name\":\"zhangsan\"}"
	//jsonStr := ""
	info := new(Info)
	err := json.Unmarshal([]byte(jsonStr), &info)
	if err != nil {
		fmt.Println(err.Error())
	}
	var in *Info
	return in
}
