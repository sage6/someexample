package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		has := strings.HasPrefix(url, "http")
		if !has {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s %v\n", url, err)
			continue
		}
		fmt.Printf("%s\n", b)
	}
}
