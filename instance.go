package hosts

import "github.com/mitchellh/goamz/ec2"

// Instance.
type Instance struct {
	ec2.Instance
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
