package mogilefs

import (
	"log"
)

type Client struct {
	Domain   string
	Backend  *Backend
	Readonly bool
}

func NewClient(args map[string]interface{}) (c *Client) {
	c = &Client{}
	return c.init(args)
}

func (c *Client) Reload(args map[string]interface{}) *Client {
	return c.init(args)
}

func (c *Client) init(args map[string]interface{}) *Client {
	switch t := args["Domain"].(type) {
	case string:
		c.Domain = t
	default:
		log.Fatal("`Domain` must be a string value.")
	}

	switch t := args["Readonly"].(type) {
	case bool:
		c.Readonly = t
	case nil:
		c.Readonly = false
	default:
		log.Fatal("`Readonly` must be a boolean value.")
	}

	if c.Backend == nil {
		var timeout int
		switch t := args["Timeout"].(type) {
		case int:
			timeout = t
		case nil:
			timeout = 3
		default:
			log.Fatal("`Timeout` must be an int value.")
		}

		c.Backend = NewBackend(map[string]interface{}{
			"Hosts":   args["Hosts"],
			"Timeout": timeout,
		})
	} else {
		c.Backend.Reload(map[string]interface{}{
			"Hosts": args["Hosts"],
		})
	}

	return c
}

func (c *Client) newFile() (file interface{}, err ReadonlyError) {
	if c.Readonly {
		return nil, ReadonlyError{}
	}
	return
}
