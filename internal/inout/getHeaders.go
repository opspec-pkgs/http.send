package inout

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func addHeaders(
	headersMap http.Header,
) error {
	headersJSONBytes, err := os.ReadFile(HeadersFilePath)
	if err != nil {
		return err
	}

	rawHeaders := map[string][]interface{}{}
	if err := json.Unmarshal(headersJSONBytes, &rawHeaders); err != nil {
		return err
	}

	for k, vSlice := range rawHeaders {
		for _, rawV := range vSlice {
			v, err := convertToString(rawV)
			if err != nil {
				return err
			}

			headersMap.Add(k, v)
		}
	}
	
	return nil
}

func convertToString(
	in interface{},
) (string, error) {
	if s, ok := in.(string); ok {
		return s, nil
	} else if b, ok := in.(bool); ok {
		return strconv.FormatBool(b), nil
	} else if f, ok := in.(float64); ok {
		return strconv.FormatFloat(f, 'f', -1, 64), nil
	} else if i, ok := in.(int64); ok {
		return strconv.FormatInt(i, 10), nil
	}

	return "", fmt.Errorf("%+v not a valid header value", in)
}
