package mogilefs

type Backend struct {
	Host    string
	Timeout int
}

func NewBackend(args map[string]interface{}) (b *Backend) {
	var host string
	switch checked := args["Host"].(type) {
	case string:
		host = checked
	default:
		panic("`domain` must be a string value.")
	}

	var timeout int
	switch checked := args["Timeout"].(type) {
	case int:
		timeout = checked
	default:
		panic("`timeout` must be an int value.")
	}

	b = &Backend{
		Host:    host,
		Timeout: timeout,
	}

	return
}
