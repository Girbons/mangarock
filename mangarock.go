package mangarock

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIEndpoint is the mangarock api endpoint :)
const APIEndpoint = "https://api.mangarockhd.com"

// MangaRockInfo is used to parse a response from the `info` endpoint
// Contains the information related to a Manga
type MangaRockInfo struct {
	Code int   `json:"code"`
	Data Manga `json:"data"`
}

// MangaRockPages is used to parse a response from `pages` endpoint
// Contains the link to the images related to a manga
type MangaRockPages struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

// Client contains only the `client` by now
// Maybe in future it can contains an ApiKey
type Client struct {
	Client *http.Client
}

// NewClient returns a Client instance
func NewClient() *Client {
	return &Client{
		Client: &http.Client{},
	}
}

// getJson decode the response body to given struct
func getJSON(response *http.Response, target interface{}) error {
	return json.NewDecoder(response.Body).Decode(target)
}

// Info returns the info related to a manga
func (m *Client) Info(comicName string) (*MangaRockInfo, error) {
	mangaInfo := new(MangaRockInfo)
	response, err := m.Client.Get(fmt.Sprintf("%s/query/web400/info?oid=%s", APIEndpoint, comicName))

	if err != nil {
		return mangaInfo, err
	}

	defer response.Body.Close()
	jsonErr := getJSON(response, &mangaInfo)

	return mangaInfo, jsonErr
}

// Pages returns the pages related to a manga
func (m *Client) Pages(comicName string) (*MangaRockPages, error) {
	mangaPages := new(MangaRockPages)

	response, err := m.Client.Get(fmt.Sprintf("%s/query/web400/pages?oid=%s", APIEndpoint, comicName))

	if err != nil {
		return mangaPages, err
	}

	defer response.Body.Close()
	jsonErr := getJSON(response, &mangaPages)

	return mangaPages, jsonErr
}
