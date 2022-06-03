package ch02

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d",
				inputs[i], expected[i], ret)
		}
	}
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	//t.Error("Error") // 测试继续
	fmt.Println("End")
}

func TestFailErrorInCode(t *testing.T) {
	fmt.Println("Start")
	//t.Fatal("Error") // 测试中止
	fmt.Println("End")
}
