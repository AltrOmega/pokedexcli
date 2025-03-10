package pokeAPI

import (
	"fmt"
	"io"
	"log"
	"net/http"
  "encoding/json"
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


//change log fatals to a err return
func getJsAsBytes(link string) []byte{
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
  return body
}

func parseLocationAreaJson(bytes []byte) EnumeratedResp {
  resp := EnumeratedResp{}
  err := json.Unmarshal(bytes, &resp)
  if err != nil {
    fmt.Println(err) 
  }
    return resp
}


func GetLocationArea(link string) EnumeratedResp{
  jsBytes := getJsAsBytes(link)
  enumResp := parseLocationAreaJson(jsBytes)
  return enumResp
}
