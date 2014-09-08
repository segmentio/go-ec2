//
// High-level EC2 instance querying library.
//
package hosts

import "github.com/mitchellh/goamz/aws"
import "github.com/mitchellh/goamz/ec2"

// Client provides some higher level methods
// for filtering, however if you need more
// flexibility you should use .Query().
type Client struct {
	ec2 *ec2.EC2
}

// New client with the given auth and region.
func New(auth aws.Auth, region aws.Region) *Client {
	return &Client{
		ec2: ec2.New(auth, region),
	}
}

// All returns all instances.
func (c *Client) All() ([]Instance, error) {
	return c.Find().Done()
}

// Running instances.
func (c *Client) Running() ([]Instance, error) {
	return c.Find().State("running").Done()
}

// Pending instances.
func (c *Client) Pending() ([]Instance, error) {
	return c.Find().State("pending").Done()
}

// Terminated instances.
func (c *Client) Terminated() ([]Instance, error) {
	return c.Find().State("terminated").Done()
}

// Stopped instances.
func (c *Client) Stopped() ([]Instance, error) {
	return c.Find().State("stopped").Done()
}

// Id running instances with the given id(s).
func (c *Client) Id(id ...string) ([]Instance, error) {
	return c.Find().State("running").Id(id...).Done()
}

// Name running instances with the given name(s).
func (c *Client) Name(name ...string) ([]Instance, error) {
	return c.Find().State("running").Name(name...).Done()
}

// Tag running  instances with the given tag and value(s).
func (c *Client) Tag(name string, value ...string) ([]Instance, error) {
	return c.Find().State("running").Tag(name, value...).Done()
}

// Find returns a chainable query.
func (c *Client) Find() *Query {
	return &Query{
		client: c,
		filter: ec2.NewFilter(),
	}
}
