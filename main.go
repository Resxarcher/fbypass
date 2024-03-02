package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type yamlData struct {
	Method  string `yaml:"method"`
	Headers string `yaml:"headers"`
}

func bypass(d string) {
	// req, err := http.NewRequest()
	// client := &http.Client{}

	yamlMap := make(map[string]interface{})
	yamlFile, err := os.ReadFile("/home/hunter/.config/fbypass/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, yamlMap)
	if err != nil {
		fmt.Println("err during unmarshaling $HOME/.config/fbypass/config.yaml:\n", err)
		os.Exit(0)
	}
	for _, item := range yamlMap {
		fmt.Println(item)
	}
}

func main() {
	var (
		host        string
		hosts       string
		concurrency int
	)
	flag.StringVar(&host, "host", "", "single host to check")
	flag.StringVar(&hosts, "hosts", "", "file path contains hosts to check")
	flag.IntVar(&concurrency, "c", 10, "number of concurrency")
	flag.Parse()
	if host != "" && hosts == "" {
		bypass(host)
	}

	if hosts == "" && host == "" {
		fmt.Println("you must provide a host or list of hosts.")
		flag.Usage()
		os.Exit(0)
	}
}
