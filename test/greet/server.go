// server.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// RouteはこのAPIサーバのルーティング設定をしている
func Route() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		// nameを受け取ってパラメータとして埋め込んで返す
		fmt.Fprintf(w, "Hello, %s", r.FormValue("name"))
	})
	return m
}

func main() {
	m := Route()
	log.Fatal(http.ListenAndServe(":8080", m))
}
