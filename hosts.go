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
func (c *Client) All() ([]ec2.Instance, error) {
	return c.Query().Get()
}

// Running instances.
func (c *Client) Running() ([]ec2.Instance, error) {
	return c.Query().State("running").Get()
}

// Pending instances.
func (c *Client) Pending() ([]ec2.Instance, error) {
	return c.Query().State("pending").Get()
}

// Terminated instances.
func (c *Client) Terminated() ([]ec2.Instance, error) {
	return c.Query().State("terminated").Get()
}

// Stopped instances.
func (c *Client) Stopped() ([]ec2.Instance, error) {
	return c.Query().State("stopped").Get()
}

// Id running instances with the given id(s).
func (c *Client) Id(id ...string) ([]ec2.Instance, error) {
	return c.Query().State("running").Id(id...).Get()
}

// Name running instances with the given name(s).
func (c *Client) Name(name ...string) ([]ec2.Instance, error) {
	return c.Query().State("running").Name(name...).Get()
}

// Tag running  instances with the given tag and value(s).
func (c *Client) Tag(name string, value ...string) ([]ec2.Instance, error) {
	return c.Query().State("running").Tag(name, value...).Get()
}

// Query returns a chainable query.
func (c *Client) Query() *Query {
	return &Query{
		client: c,
		filter: ec2.NewFilter(),
	}
}
