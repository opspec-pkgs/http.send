package inout

import (
	"net/http"
	"os"
	"strconv"
)

func SetResponse(
	resp *http.Response,
) error {

	if err := setBody(resp.Body); err != nil {
		return err
	}
	if err := setHeaders(resp.Header); err != nil {
		return err
	}
	if err := os.WriteFile(ProtocolFilePath, []byte(resp.Proto), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(StatusCodeFilePath, []byte(strconv.Itoa(resp.StatusCode)), 0644); err != nil {
		return err
	}

	return nil
}
