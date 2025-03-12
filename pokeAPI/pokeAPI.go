package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pokedexcli/pokecache"
	"time"
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

var pcache *pokecache.Cache

func getJsAsBytes(link string) ([]byte, error) {
	// kinda meh solution would be better to check it only once
	// and not on every get call
	if pcache == nil {
		pcache = pokecache.NewCache(time.Second * 5)
	}
	val, ok := pcache.Get(link)
	if ok {
		return val, nil
	}

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
	pcache.Add(link, body)
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
