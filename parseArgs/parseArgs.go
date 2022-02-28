package parseArgs

import (
	"os"
	"strings"

	"github.com/roceb/gUrl/request"
)

func ArgParser() []string {
	args := os.Args[1:]
	firstArg := strings.ToUpper(args[0])
	var req []string
	if firstArg != "GET" && firstArg != "POST" && firstArg != "PUT" && firstArg != "DELETE" {
		req = append(req, "GET")
	}
	switch len(args) {
	case 1:
		req = append(req, "GET")
		req = append(req, args[0], "", "")
	case 2:
		req = append(req, strings.ToUpper(args[0]), args[1], "", "")
	case 3:
		req = append(req, strings.ToUpper(args[0]), args[1], args[2], "")
	default:
		req = append(req, args...)
	}
	return req
}

func AssignVerb(c *request.ParsedArgs, v string) {
	c.Verb = v
}
func AssignUrl(c *request.ParsedArgs, u string) {
	if u == "" {
		panic("No url was passed")
	}
	c.Url = u
}

func AssignData(c *request.ParsedArgs, d []string) {
	data := make(map[string]interface{})
	for _, v := range d {
		i := strings.Index(v, "=")
		key := v[:i]
		value := v[i+1:]
		data[key] = value
	}
	c.Data = data
}
