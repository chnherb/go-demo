package main

import "fmt"

func main() {
	//strBasic()
	//strUnicode()
	strToRune()
}

func strBasic() {
	var s string
	fmt.Printf("s:%s*\n", s)
	s = "hello"
	fmt.Println(len(s)) // 5
	// s[1] = '3' // error, string 是不可变的 byte slice
	s = "\xE4\xB8\xA5" // 可以存储任意二进制数据
	fmt.Printf("s:%s*\n", s) // s:严*
	fmt.Println(len(s)) // 3  byte数
}

func strUnicode() {
	s := "中"
	fmt.Printf("s:%s*\n", s) // s:中*
	fmt.Println(len(s)) // 3  byte数

	c := []rune(s) // 取出字符串的 Unicode
	fmt.Println(len(c)) // 1

	fmt.Printf("中 Unicode %x\n", c[0]) // 中 Unicode 4e2d
	fmt.Printf("中 UTF8 %x\n", s) // 中 UTF8 e4b8ad
}

func strToRune() {
	s := "中华人民共和国"
	for _, c := range s {
		fmt.Printf("%[1]c, %[1]x\n", c)
	}
}