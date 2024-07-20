package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var ctxBg = context.Background()
var redisClient *redis.Client;

func main()  {
	fmt.Println("A Sample Web Server with Redis Caching")

	// initialize the redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	http.HandleFunc("/api/devrelData", getData) //defining the api endpoint
	http.ListenAndServe(":8080", nil) //defining the port
}

func getData(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello fellow DevRel's")

	// code to fetch the data
}