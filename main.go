package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	_ "gorm_test/setting"
	"net/http"
	"strconv"
)

func test() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		fmt.Println("path:", r.URL.Path)
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])

		w.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		w.Write(jData)
	})
	http.ListenAndServe(":8000", nil)
}

func main() {
	defer zap.L().Sync()
	test()
}
