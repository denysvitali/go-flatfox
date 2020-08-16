package flatfox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const ApiEndpoint = "https://flatfox.ch/api"

type Client struct {
	client *http.Client
}

func NewClient(client *http.Client) *Client {
	return &Client{
		client: client,
	}
}

type Bounds struct {
	North float64
	West  float64
	South float64
	East  float64
}

type PinSearchOptions struct {
	MinRooms       *float32        `json:"min_rooms"`
	MaxRooms       *float32        `json:"max_rooms"`
	MinSpace       *int32          `json:"min_space"`
	MaxSpace       *int32          `json:"max_space"`
	MinFloor       *int32          `json:"min_floor"`
	MaxFloor       *int32          `json:"max_floor"`
	MinPrice       *int32          `json:"min_price"`
	MaxPrice       *int32          `json:"max_price"`
	MinYearBuilt   *int32          `json:"min_year_built"`
	MaxYearBuilt   *int32          `json:"max_year_built"`
	IsFurnished    *bool           `json:"is_furnished"`
	IsTemporary    *bool           `json:"is_temporary"`
	ObjectCategory *ObjectCategory `json:"object_category"`
}

func (c *Client) PinSearch(count int, bounds *Bounds, searchOptions *PinSearchOptions) (*[]Flat, error) {

	baseUrl, err := url.Parse(fmt.Sprintf("%s/v1/pin", ApiEndpoint))
	if err != nil {
		return nil, err
	}

	baseUrl.Query().Add("count", strconv.Itoa(count))
	baseUrl.Query().Add("north", fmt.Sprintf("%f", bounds.North))
	baseUrl.Query().Add("south", fmt.Sprintf("%f", bounds.South))
	baseUrl.Query().Add("west", fmt.Sprintf("%f", bounds.West))
	baseUrl.Query().Add("east", fmt.Sprintf("%f", bounds.East))

	if searchOptions.MinRooms != nil {
		baseUrl.Query().Add("min_rooms", fmt.Sprintf("%f", *searchOptions.MinRooms))
	}

	if searchOptions.MaxRooms != nil {
		baseUrl.Query().Add("min_rooms", fmt.Sprintf("%f", *searchOptions.MaxRooms))
	}

	req, err := http.NewRequest(http.MethodGet, baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	var flats []Flat
	var jsonDecoder = json.NewDecoder(resp.Body)
	jsonDecoder.DisallowUnknownFields()


	err = jsonDecoder.Decode(&flats)
	if err != nil {
		return nil, err
	}

	return &flats, nil
}
