package main

import "github.com/mitchellh/goamz/aws"
import "github.com/segmentio/go-hosts"
import "encoding/json"
import "os"

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

	nodes, err := hosts.All()
	check(err)

	b, err := json.MarshalIndent(nodes, "", "  ")
	check(err)

	os.Stdout.Write(b)
}
