package main

import "fmt"
import "net/http"

func main() {
	fmt.Printf("Hello, world.\n")
	resp, err := http.Get("http://example.com/")

	fmt.Printf(resp)
}
