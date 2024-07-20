package main

import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("A Sample Web Server with Redis Caching")

	http.HandleFunc("/api/devrelData", getData)
	http.ListenAndServe(":8080", nil)
}

func getData(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello fellow DevRel's")
}