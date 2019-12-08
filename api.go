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

// NewClient Creates a new WizzClient
func NewClient() (*WizzClient, error) {
	return NewCustomClient(MetadataURL)
}

// NewCustomClient Creates a new WizzClient using provided matadata URL to retrieve current Wizzair API URL
func NewCustomClient(metadataURL string) (*WizzClient, error) {
	httpClient := resty.New()
	if resp, err := httpClient.R().Get(metadataURL); err != nil {
		return nil, err
	} else if metadataDto, err := parseMetadataDto(resp.Body()); err != nil {
		return nil, err
	} else {
		return &WizzClient{
			client: httpClient.
				SetHostURL(metadataDto.ApiURL).
				SetCloseConnection(true).
				SetDebug(false),
		}, nil
	}
}

// CopyClient Creates copy of existing client.
// Motivation: Using same client to do 2+ request causes 2nd request to be rejected by Wizzair
// (could not figured out why, closing connection didn't helped, cleaning cookies no luck either)
func CopyClient(wizz *WizzClient) *WizzClient {
	httpClient := resty.New()
	return &WizzClient{
		client: httpClient.
			SetHostURL(wizz.client.HostURL).
			SetCloseConnection(true).
			SetDebug(false),
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
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", UserAgent).
		// SetHeader("Cookie", "").
		SetBody(reqDto).
		SetResult(respDto).
		Post(path)

	if resp != nil && resp.IsError() {
		return fmt.Errorf("Request failed. Status %d. Body: %s", resp.StatusCode(), string(resp.Body()))
	}

	return err
}

func parseMetadataDto(data []byte) (*MetadataDto, error) {
	r := &MetadataDto{}
	return r, json.Unmarshal(data, r)
}
