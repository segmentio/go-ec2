
# go-hosts

 High-level EC2 instance querying library built on top of [mitchellh/goamz/ec2](https://github.com/mitchellh/goamz/tree/master/ec2).

 View the [docs](http://godoc.org/github.com/segmentio/go-hosts).

## Example

```go
auth, err := aws.EnvAuth()
check(err)

hosts := hosts.New(auth, aws.USWest2)
check(err)

nodes, err := hosts.Name("api-*")
check(err)

for i, node := range nodes {
  fmt.Printf("  %d) %s\n", i, node.InstanceId)
}
```

# License

 MIT
