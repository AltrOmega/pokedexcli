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

var pcache *pokecache.Cache

func getJsAsBytes(link string) ([]byte, error) {
	if pcache == nil {
		pcache = pokecache.NewCache(time.Second * 60)
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

func GetLink(endpoint, destination string) string {
	return fmt.Sprintf("%v%v", endpoint, destination)
}

func parseJson[T any](bytes []byte) (T, error) {
	var resp T
	err := json.Unmarshal(bytes, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func GetResp[T any](link string) (T, error) {
	var resp T
	jsBytes, err := getJsAsBytes(link)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	resp, err = parseJson[T](jsBytes)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	return resp, nil
}
