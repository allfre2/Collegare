package util

import (
	"bytes"
	"encoding/json"
	"strings"
)

func Serialize[T any](obj *T) (string, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	if err := encoder.Encode(*obj); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func Deserialize[T any](input string) (*T, error) {
	stream := strings.NewReader(input)
	decoder := json.NewDecoder(stream)
	obj := new(T)
	if err := decoder.Decode(obj); err != nil {
		return nil, err
	}

	return obj, nil
}
