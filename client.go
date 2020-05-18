package gowizz

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// WizzClient HTTP client to get data from Wizzair
type WizzClient struct {
	client *resty.Client
}

type parser func([]byte) (interface{}, error)

// NewClientOrErr Creates a new WizzClient or panic with error
func NewClientOrErr(debug bool) *WizzClient {
	if client, err := NewCustomClient(MetadataURL, debug); err != nil {
		panic(err)
	} else {
		return client
	}
}

// NewClient Creates a new WizzClient
func NewClient(debug bool) (*WizzClient, error) {
	return NewCustomClient(MetadataURL, debug)
}

// NewCustomClient Creates a new WizzClient using provided matadata URL to retrieve current Wizzair API URL
func NewCustomClient(metadataURL string, debug bool) (*WizzClient, error) {
	httpClient := resty.New().SetDebug(debug)
	if resp, err := httpClient.R().SetHeader("User-Agent", UserAgent).Get(metadataURL); err != nil {
		return nil, err
	} else if metadataDto, err := parseMetadataDto(resp.Body()); err != nil {
		return nil, err
	} else {
		return &WizzClient{
			client: httpClient.
				SetHostURL(metadataDto.ApiURL).
				SetCloseConnection(false),
		}, nil
	}
}

// SearchFlights Search flights
func (wizz *WizzClient) SearchFlights(filter SearchFilterDto) (*SearchResultDto, error) {
	respDto := &SearchResultDto{}
	return respDto, wizz.doPost(SearchPath, filter, respDto)
}

// TimetableSearch Search timetable
func (wizz *WizzClient) TimetableSearch(filter TimetableSearchFilterDto) (*TimetableSearchResultDto, error) {
	respDto := &TimetableSearchResultDto{}
	return respDto, wizz.doPost(TimetableSearchPath, filter, respDto)
}

// GetCities Get information about cities Wizzair flight to/from
func (wizz *WizzClient) GetCities() (*CitiesDto, error) {
	bodyDto := &CitiesDto{}
	return bodyDto, wizz.doGet(CitiesPath, bodyDto)
}

func (wizz *WizzClient) doGet(path string, respDto interface{}) error {
	resp, err := wizz.client.R().
		SetResult(respDto).
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetHeader("User-Agent", UserAgent).
		SetHeader("Accept-Encoding", "gzip, deflate, br").
		Get(path)

	if resp != nil && resp.IsError() {
		return fmt.Errorf("Request failed. Status %d. Body: %s", resp.StatusCode(), string(resp.Body()))
	}

	return err
}

func (wizz *WizzClient) doPost(path string, reqDto interface{}, respDto interface{}) error {
	resp, err := wizz.client.
		// EnableTrace().
		NewRequest().
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetHeader("User-Agent", UserAgent).
		SetHeader("Accept-Encoding", "gzip, deflate, br").
		SetBody(reqDto).
		SetResult(respDto).
		Post(path)

	// fmt.Print(string(resp.Body()))

	if resp != nil && resp.IsError() {
		return fmt.Errorf("Request failed. Status %d. Body: %s", resp.StatusCode(), string(resp.Body()))
	}

	return err
}

func parseMetadataDto(data []byte) (*MetadataDto, error) {
	r := &MetadataDto{}
	return r, json.Unmarshal(data, r)
}
