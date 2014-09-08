package hosts

import "github.com/mitchellh/goamz/aws"
import "github.com/mitchellh/goamz/ec2"

// Client.
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
	return c.Query().Get()
}

// Running instances.
func (c *Client) Running() ([]Instance, error) {
	return c.Query().State("running").Get()
}

// Pending instances.
func (c *Client) Pending() ([]Instance, error) {
	return c.Query().State("pending").Get()
}

// Terminated instances.
func (c *Client) Terminated() ([]Instance, error) {
	return c.Query().State("terminated").Get()
}

// Stopped instances.
func (c *Client) Stopped() ([]Instance, error) {
	return c.Query().State("stopped").Get()
}

// Id running instances with the given id(s).
func (c *Client) Id(id ...string) ([]Instance, error) {
	return c.Query().State("running").Id(id...).Get()
}

// Name running instances with the given name(s).
func (c *Client) Name(name ...string) ([]Instance, error) {
	return c.Query().State("running").Name(name...).Get()
}

// Tag running  instances with the given tag and value(s).
func (c *Client) Tag(name string, value ...string) ([]Instance, error) {
	return c.Query().State("running").Tag(name, value...).Get()
}

// Query returns a chainable query.
func (c *Client) Query() *Query {
	return &Query{
		client: c,
		filter: ec2.NewFilter(),
	}
}
