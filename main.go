package main

import (
	"fmt"
	"net/http"

	"github.com/sem-onyalo/application-dashboard/model"
)

func main() {
	ep := model.Endpoint{Name: "Endpoint 1", URL: "http://api.wazomill.com"}
	rsp, err := http.Get(ep.URL)
	var result string

	if err != nil {
		result = "FAIL"
	} else {
		result = "PASS"
	}

	fmt.Printf("%s %s %s\n", ep.Name, result, rsp.Status)
}
