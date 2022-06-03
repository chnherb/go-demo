package ch02

import (
	"bytes"
	"testing"
)

func BenchmarkDemo(b *testing.B) {
	// 与性能测试无关的代码
	elems := []string{"1","2","3", "4", "5"}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		var buf bytes.Buffer
		for _, elem := range elems{
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
