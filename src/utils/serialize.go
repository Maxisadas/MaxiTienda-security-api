package utils

import (
	"api-rest/src/errors"
	"encoding/json"
	"io"
)

func Deserialize(body io.ReadCloser, obj interface{}) *errors.HTTPErrors {
	err := json.NewDecoder(body).Decode(&obj)
	if err != nil {
		var e = errors.HTTPErrors {
			Message: err.Error(),
			Code: 500,
		}
		return &e
	}
	defer body.Close()
	return nil
}