
# go-ec2

 High-level EC2 library built on top of [mitchellh/goamz/ec2](https://github.com/mitchellh/goamz/tree/master/ec2).

 View the [docs](http://godoc.org/github.com/segmentio/go-ec2).

## Features

 Currently just provides a fluent API for looking up instances.

## Example

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

# License

 MIT
