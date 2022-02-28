package response

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/roceb/gUrl/request"
)

func PreRequest(c *request.ParsedArgs) *http.Request {
	req, err := http.NewRequest(c.Verb, c.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	return req
}

// Calls a newRequest and takes in a ParsedArgs
func Call(c *request.ParsedArgs, req *http.Request) string {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(f)

}
