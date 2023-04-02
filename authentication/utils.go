package main

import (
	"encoding/json"
	"io"
)

func decodeBody[T any](reader io.Reader) (*T, error) {
	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var dest T
	err = json.Unmarshal(bodyBytes, &dest)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}
