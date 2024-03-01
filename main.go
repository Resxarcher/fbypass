package main

import (
	"flag"
)

func bypass(d string) {

}

func main() {
	var (
		domain      string
		lists       string
		concurrency int
	)
	flag.StringVar(&domain, "domain", "", "single domain to check")
	flag.StringVar(&lists, "lists", "", "file path contains domains to check")
	flag.IntVar(&concurrency, "c", 10, "number of concurrency")
	flag.Parse()

}
