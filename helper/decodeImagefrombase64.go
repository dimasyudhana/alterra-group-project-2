package helper

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func DecodeImageFromBase64(encodedString string, filename string) error {
	decoded, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, decoded, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
