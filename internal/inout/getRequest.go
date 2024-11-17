package inout

import (
	"net/http"
	"os"
)

func GetRequest() (
	*http.Request,
	error,
) {
	methodBytes, err := os.ReadFile(MethodFilePath)
	if err != nil {
		return nil, err
	}

	urlBytes, err := os.ReadFile(URLFilePath)
	if err != nil {
		return nil, err
	}

	reqBody, err := os.Open(BodyFilePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		string(methodBytes),
		string(urlBytes),
		reqBody,
	)
	if err != nil {
		return nil, err
	}

	reqBodyStat, err := reqBody.Stat()
	if err != nil {
		return nil, err
	}

	req.ContentLength = reqBodyStat.Size()

	if err := addHeaders(req.Header); err != nil {
		return nil, err
	}

	return req, nil
}
