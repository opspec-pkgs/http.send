package inout

import (
	"io"
	"os"
)

func setBody(
	body io.ReadCloser,
) error {
	bodyFile, err := os.Create(BodyFilePath)
	if err != nil {
		return err
	}
	defer bodyFile.Close()

	_, err = io.Copy(bodyFile, body)
	defer body.Close()
	return err
}
