package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// https://pokeapi.co/api/v2/location-area/

type EnumeratedResp struct {
	Count    int          `json:"count"`
	Next     *string      `json:"next"`
	Previous *string      `json:"previous"`
	Results  []NameAndUrl `json:"results"`
}

type NameAndUrl struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// change log fatals to a err return
func getJsAsBytes(link string) ([]byte, error) {
	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

func parseLocationAreaJson(bytes []byte) (EnumeratedResp, error) {
	resp := EnumeratedResp{}
	err := json.Unmarshal(bytes, &resp)
	if err != nil {
		fmt.Println(err)
		return EnumeratedResp{}, err
	}
	return resp, nil
}

func GetLocationArea(link string) (EnumeratedResp, error) {
	jsBytes, err := getJsAsBytes(link)
	if err != nil {
		fmt.Println(err)
		return EnumeratedResp{}, err
	}

	enumResp, err := parseLocationAreaJson(jsBytes)
	if err != nil {
		fmt.Println(err)
		return EnumeratedResp{}, err
	}

	return enumResp, nil
}
