package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctxBg = context.Background()
var redisClient *redis.Client;

func main()  {
	fmt.Println("A Sample Web Server with Redis Caching")

	// initialize the redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:<PORT_NUMBER>",
		Password: "",
		DB: 0,
	})

	http.HandleFunc("/api/devrelData", getData) //defining the api endpoint
	http.ListenAndServe(":8080", nil) //defining the port
}

func getData(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello fellow DevRel's")

	// code to fetch the data
	cacheKey := "api-data-cache"

	// check if cache key exists in Redis instance
	cachedData, err := redisClient.Get(ctxBg, cacheKey).Result()

	// error handling
	if err == redis.Nil{
		// if cache key not found, fetch and store the data in Redis
		data := "Hello Aman"
		cacheExpiration := 10*time.Minute

		err := redisClient.Set(ctxBg, cacheKey, data, cacheExpiration).Err()

		if err!=nil {
			panic(err)
		}

		fmt.Fprintf(w, data)
	}else if err!=nil {
		panic(err)
	}else {
		// cache key found, hence return this data
		fmt.Fprintf(w, cachedData)
	}
}