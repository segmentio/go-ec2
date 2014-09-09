package ec2

import "github.com/mitchellh/goamz/ec2"

// Query provides a chainable API
// for building up filters.
type Query struct {
	client *Client
	ids    []string
	filter *ec2.Filter
}

// Id filters on instance id(s).
func (q *Query) Id(id ...string) *Query {
	q.ids = append(q.ids, id...)
	return q
}

// AvailabilityZone filters on AZ.
func (q *Query) AvailabilityZone(name string) *Query {
	q.filter.Add("availability-zone", name)
	return q
}

// ImageId filters on the image used to launch the instance.
func (q *Query) ImageId(id string) *Query {
	q.filter.Add("image-id", id)
	return q
}

// Type filters on the instance type, for example "m1.small".
func (q *Query) Type(name string) *Query {
	q.filter.Add("instance-type", name)
	return q
}

// Spot filters spot instances.
func (q *Query) Spot() *Query {
	q.filter.Add("instance-lifecycle", "spot")
	return q
}

// PublicIp filters on public ip address.
func (q *Query) PublicIp(ip string) *Query {
	q.filter.Add("ip-address", ip)
	return q
}

// PrivateIp filters ip address.
func (q *Query) PrivateIp(ip string) *Query {
	q.filter.Add("private-ip-address", ip)
	return q
}

// KeyName filters key-pair name.
func (q *Query) KeyName(name string) *Query {
	q.filter.Add("key-name", name)
	return q
}

// State filters on the instance state.
func (q *Query) State(state string) *Query {
	q.filter.Add("instance-state-name", state)
	return q
}

// Tag filters on tag name and value.
func (q *Query) Tag(name string, value ...string) *Query {
	q.filter.Add("tag:"+name, value...)
	return q
}

// Name filters on the Name tag.
func (q *Query) Name(name ...string) *Query {
	return q.Tag("Name", name...)
}

// Done kicks off the request and returns matching instance(s).
func (q *Query) Done() ([]Instance, error) {
	if res, err := q.client.ec2.Instances(q.ids, q.filter); err == nil {
		return flatten(res), nil
	} else {
		return nil, err
	}
}

// flat list of instances sans reservations.
func flatten(res *ec2.InstancesResp) []Instance {
	var instances []Instance
	for _, reservation := range res.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, Instance{instance})
		}
	}
	return instances
}
