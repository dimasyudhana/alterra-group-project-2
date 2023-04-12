package helper

import (
	"encoding/base64"
)

func EncodeImageToBase64(image []byte) (string, error) {
	encodedString := base64.StdEncoding.EncodeToString(image)
	return encodedString, nil
}
