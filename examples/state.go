package main

import "github.com/mitchellh/goamz/aws"
import "github.com/segmentio/go-ec2"
import "fmt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	auth, err := aws.EnvAuth()
	check(err)

	hosts := ec2.New(auth, aws.USWest2)
	check(err)

	{
		nodes, err := hosts.Running()
		check(err)
		fmt.Printf("%d nodes running\n", len(nodes))
	}

	{
		nodes, err := hosts.Stopped()
		check(err)
		fmt.Printf("%d nodes stopped\n", len(nodes))
	}
}
