package hosts

import "github.com/mitchellh/goamz/ec2"

// flat list of instances sans reservations.
func instances(res *ec2.InstancesResp) []Instance {
	var instances []Instance
	for _, reservation := range res.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, Instance{instance})
		}
	}
	return instances
}
