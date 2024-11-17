package main

import (
	"net/http"

	"github.com/opspec-pkgs/http.send/internal/inout"
)

func main() {

	req, err := inout.GetRequest()
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	if err := inout.SetResponse(resp); err != nil {
		panic(err)
	}

}