package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"mvdan.cc/xurls"
)

// https://stackoverflow.com/a/20895629
// https://godoc.org/?q=regexp
// https://github.com/mvdan/xurls

var (
	link    = flag.String("l", "http://google.com", "The first link")
	linkNum = flag.Int("n", 100, "Number of links")
)

func main() {
	flag.Parse()

	candidates := []string{*link}
	all := []string{}

	for len(candidates) > 0 && len(all) < *linkNum {
		url := candidates[0]
		candidates = candidates[1:]
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("http.Get: %v", err)
		}
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("iotil.ReaAll: %v", err)
		}
		resultlinks := xurls.Strict().FindAllString(string(bytes), *linkNum)
		for _, url := range resultlinks {
			candidates = append(candidates, url)
			all = append(all, url)
		}
	}

	if len(all) > *linkNum {
		all = all[:*linkNum]
	}

	for _, result := range all {
		fmt.Println(result)
	}
}
