package main

import (
	"fmt"
	"net/http"
)

func main() {
	accessAction("/")
	accessAction("/test")
	accessAction("/healthz")
}

func accessAction(path_name string) {
	resp, err := http.Get("http://192.168.50.184:8090" + path_name)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response: ", resp.Header)
}
