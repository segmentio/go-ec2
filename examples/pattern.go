package main

import "github.com/mitchellh/goamz/aws"
import "github.com/segmentio/go-hosts"
import "fmt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	auth, err := aws.EnvAuth()
	check(err)

	hosts := hosts.New(auth, aws.USWest2)
	check(err)

	nodes, err := hosts.Name("api-*")
	check(err)

	for i, node := range nodes {
		fmt.Printf("  %d) %s\n", i, node.Tag("Name"))
	}
}
