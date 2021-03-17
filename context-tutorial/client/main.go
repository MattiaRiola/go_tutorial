package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		fmt.Println("enter 'disc' to stop the client:")
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		for s.Text() != "disc" {
			fmt.Println("wrong command")
			s.Scan()
		}
		cancel()
	}()

	//create a new request req
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	//add the context in the request req and it return my new req with the ctx inside
	req = req.WithContext(ctx)
	//send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}
