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
	log.Printf("handler started")
	defer log.Printf("handelr ended")
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "hello")
}
