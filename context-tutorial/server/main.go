package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("handler started")
	defer log.Printf("handelr ended")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		//in this way if the client will interrupt the connection I receive the ctx.Done()
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
