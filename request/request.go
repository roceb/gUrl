package request

import (
	"net/http"
	"strings"
)

type ParsedArgs struct {
	Verb   string
	Url    string
	Data   map[string]interface{}
	Header *http.Request
}

func AddHeader(req *http.Request, h []string) {
	req.Header.Set("User-Agent", "gUrl/1.1")
	for _, v := range h {
		i := strings.Index(v, ":")
		key := v[:i]
		value := v[i+1:]
		req.Header.Set(key, value)
	}
}
