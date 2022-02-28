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
	for _, v := range h {
		i := strings.Index(v, ":")
		key := v[:i]
		value := v[i+1:]
		req.Header.Set(key, value)
		// fmt.Println("K V ", key, value)

		// fmt.Println(c.Header)
	}
}

// type Header map[string][]string

// func (h *Header) Add(key, value string) {
// 	// (*h)[key] = append((*h)[key], value)
// 	var a []string
// 	(*h)[key] = append(a, value)
// 	fmt.Print((h))
// }
// func (h *Header) Get(v string) []string {
// 	ans := (*h)[v]
// 	return ans
// }
