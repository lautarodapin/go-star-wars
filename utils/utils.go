package utils

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
	Detail string `json:"detail"`
}

type ApiResponse[T any] struct {
	Count    int32   `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []T     `json:"results"`
}

func MakeRequest[T any](url string, data *T) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&data)
	return nil
}
