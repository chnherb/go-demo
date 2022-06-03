package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// /time 后面如果没有 / 则不会匹配到子路径
	// 如 /time/ 和 /time/xx 都会匹配到 /

	// /time/ 下
	// /time/ 和 /time/xx 都会匹配到此路径下
	// 注意此时 /time 在浏览器下回自动补齐为 /time/，但curl xxx/time 则不会匹配到
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}
