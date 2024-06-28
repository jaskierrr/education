package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var (
	ctx = context.Background()
	//rdb *redis.Client
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
})

func SetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("Unadble to set data in redis: %v", err)
	}
	fmt.Fprintf(w, "Key %s set to %s\n", key, value)
}

func GetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Unadble to set data in redis: %v", err)
	}

	url := "https://"+value

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
