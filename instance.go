package hosts

import "github.com/mitchellh/goamz/ec2"

// Instance.
type Instance struct {
	ec2.Instance
}

func (i *Instance) Tag(name string) string {
	for _, tag := range i.Tags {
		if tag.Key == name {
			return tag.Value
		}
	}
	return ""
}
