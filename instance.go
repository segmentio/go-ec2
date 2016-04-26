package ec2

import "github.com/mitchellh/goamz/ec2"

// Instance.
type Instance struct {
	ec2.Instance
}

// ID returns the instance ID.
func (i *Instance) ID() string {
	return i.InstanceId
}

// Tag returns the tag value of `name` or an empty string.
func (i *Instance) Tag(name string) string {
	for _, tag := range i.Tags {
		if tag.Key == name {
			return tag.Value
		}
	}
	return ""
}

// Name returns the "Name" tag value.
func (i *Instance) Name() string {
	return i.Tag("Name")
}
