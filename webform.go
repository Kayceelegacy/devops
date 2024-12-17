package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	color := os.Getenv("BG_COLOR")
	if color == "" {
		color = "blue" // default  color if not specified
	}
	fmt.Fprintf(w, "<body style='background-color:%s'><h1 style='font-size:70px;color:pink'>Hello, %s!</h1></body>", color, name)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

