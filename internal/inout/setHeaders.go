package inout

import (
	"encoding/json"
	"net/http"
	"os"
)

func setHeaders(
	headers http.Header,
) error {
	headersJSONBytes, err := json.Marshal(headers)
	if err != nil {
		return err
	}

	return os.WriteFile(
		HeadersFilePath,
		headersJSONBytes,
		0644,
	)
}
