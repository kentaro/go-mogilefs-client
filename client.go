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
	var domain string
	switch t := args["Domain"].(type) {
	case string:
		domain = t
	default:
		log.Fatal("`Domain` must be a string value.")
	}

	var timeout int
	switch t := args["Timeout"].(type) {
	case int:
		timeout = t
	case nil:
		timeout = 3
	default:
		log.Fatal("`Timeout` must be an int value.")
	}

	var readonly bool
	switch t := args["Readonly"].(type) {
	case bool:
		readonly = t
	case nil:
		readonly = false
	default:
		log.Fatal("`Readonly` must be a boolean value.")
	}

	backend := NewBackend(map[string]interface{}{
		"Hosts":   args["Hosts"],
		"Timeout": timeout,
	})

	c = &Client{
		Domain:   domain,
		Backend:  backend,
		Readonly: readonly,
	}

	return
}
