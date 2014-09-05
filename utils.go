package hosts

import "github.com/mitchellh/goamz/ec2"

// flat list of instances sans reservations.
func instances(res *ec2.InstancesResp) []ec2.Instance {
	var instances []ec2.Instance
	for _, reservation := range res.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, instance)
		}
	}
	return instances
}
