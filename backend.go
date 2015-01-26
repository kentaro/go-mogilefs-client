package mogilefs

type Backend struct {
	Hosts   []string
	Timeout int
}

func NewBackend(args map[string]interface{}) (b *Backend) {
	var hosts []string
	switch t := args["Hosts"].(type) {
	case string:
		hosts = []string{t}
	case []string:
		for _, v := range t {
			hosts = append(hosts, v)
		}
	default:
		panic("`Hosts` must be either a string value or an array of strings.")
	}

	var timeout int
	switch t := args["Timeout"].(type) {
	case int:
		timeout = t
	case nil:
		timeout = 3
	default:
		panic("`Timeout` must be an int value.")
	}

	b = &Backend{
		Hosts:   hosts,
		Timeout: timeout,
	}

	return
}
