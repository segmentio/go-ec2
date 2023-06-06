
# go-ec2

> **Note**  
> Segment has paused maintenance on this project, but may return it to an active status in the future. Issues and pull requests from external contributors are not being considered, although internal contributions may appear from time to time. The project remains available under its open source license for anyone to use.

 High-level EC2 library built on top of [mitchellh/goamz/ec2](https://github.com/mitchellh/goamz/tree/master/ec2).

 View the [docs](http://godoc.org/github.com/segmentio/go-ec2).

## Features

 Currently just provides a fluent API for looking up instances.

## Example

 Lookup running nodes matching "api-*":

```go
auth, err := aws.EnvAuth()
check(err)

hosts := hosts.New(auth, aws.USWest2)
check(err)

nodes, err := hosts.Name("api-*")
check(err)

for i, node := range nodes {
  fmt.Printf("  %d) %s %s\n", i, node.InstanceId, node.Name())
}
```

 Lookup stopped nodes with the "Project" tag of "redshift":

```go
nodes, err := hosts.Find().State("stopped").Tag("Type", "redshift").Done()
```

# License

 MIT
