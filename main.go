package main

import (
	"fmt"
	"strings"

	"github.com/roceb/gUrl/parseArgs"
	"github.com/roceb/gUrl/request"
	"github.com/roceb/gUrl/response"
)

func main() {
	// Think about adding Cobra
	c := request.ParsedArgs{}
	// spits out all arguements created after
	args := parseArgs.ArgParser()
	// assign Verb GET/PUT/ ...
	parseArgs.AssignVerb(&c, strings.ToUpper(args[0]))
	// assign url
	parseArgs.AssignUrl(&c, args[1])
	extraArgs := args[2:]
	var headers []string
	var data []string
	for _, v := range extraArgs {
		if strings.Contains(v, ":") {
			headers = append(headers, v)
		}
		if strings.Contains(v, "=") {
			data = append(data, v)
		}
	}
	parseArgs.AssignData(&c, data)
	req := response.PreRequest(&c)
	request.AddHeader(req, headers)
	resp := response.Call(&c, req)

	fmt.Print(resp)
}
