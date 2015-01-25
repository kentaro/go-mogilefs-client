package mogilefs

import (
	"log"
)

type Client struct {
	Domain   string
	Backend  []*Backend
	Readonly bool
}

func NewClient(args map[string]interface{}) (c *Client) {
	var domain string
	switch checked := args["Domain"].(type) {
	case string:
		domain = checked
	default:
		log.Fatal("`domain` must be a string value.")
	}

	var timeout int
	switch checked := args["Timeout"].(type) {
	case int:
		timeout = checked
	case nil:
		timeout = 0
	default:
		log.Fatal("`timeout` must be an int value.")
	}

	var readonly bool
	switch checked := args["Readonly"].(type) {
	case bool:
		readonly = checked
	case nil:
		readonly = false
	default:
		log.Fatal("`readonly` must be a boolean value.")
	}

	var hosts []*Backend

	switch checked := args["Host"].(type) {
	case string:
		host := NewBackend(map[string]interface{}{
			"Host":    checked,
			"Timeout": timeout,
		})
		hosts = append(hosts, host)
	case []string:
		for _, v := range checked {
			host := NewBackend(map[string]interface{}{
				"Host":    v,
				"Timeout": timeout,
			})
			hosts = append(hosts, host)
		}
	default:
		log.Fatal("domain must be a string value.")
	}

	c = &Client{
		Domain:   domain,
		Backend:  hosts,
		Readonly: readonly,
	}

	return
}
