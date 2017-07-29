package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	resp, err := http.Get("http://example.com/")
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s", b)

}
