package gowizz

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	// when
	client, err := NewClient(MetadataURL)

	// then
	require.Nil(t, err)
	assert.NotNil(t, client)
}

func TestGetCities(t *testing.T) {
	client, _ := NewClient(MetadataURL)

	// when
	respDto, err := client.GetCities()

	// then
	require.Nil(t, err)
	require.NotNil(t, respDto)

	// and response contains at least 1 city
	require.NotEmpty(t, respDto.Cities)
	assert.Equal(t, respDto.Cities[0].Iata, "TIA")
	assert.Equal(t, respDto.Cities[0].CountryName, "Albania")
	assert.Equal(t, respDto.Cities[0].ShortName, "Tirana")

	// and city contains at least 1 connection
	require.NotEmpty(t, respDto.Cities[0].Connections)
	assert.Equal(t, respDto.Cities[0].Connections[0].Iata, "BUD")
}

func TestSearchFlights(t *testing.T) {
	client, _ := NewClient(MetadataURL)

	reqDto := SearchFilterDto{
		FlightList: []FlightFilter{
			FlightFilter{
				DepartureStation: "TAT",
				ArrivalStation:   "LTN",
				DepartureDate:    "2019-12-18",
			},
		},
		AdultCount:  1,
		ChildCount:  0,
		InfantCount: 0,
		Wdc:         false,
	}

	// when
	respDto, err := client.SearchFlights(reqDto)
	fmt.Printf("%+v\n", respDto)

	// then
	require.Nil(t, err)
	require.NotNil(t, respDto)
}

func TestTimetableSearch(t *testing.T) {
	client, _ := NewClient(MetadataURL)

	reqDto := TimetableSearchFilterDto{
		FlightList: []TimetableFlightFilter{
			TimetableFlightFilter{
				DepartureStation: "TAT",
				ArrivalStation:   "LTN",
				From:             time.Now().AddDate(0, 0, 31).Format("2006-01-02"),
				To:               time.Now().AddDate(0, 0, 61).Format("2006-01-02"),
			},
		},
		AdultCount:  1,
		ChildCount:  0,
		InfantCount: 0,
		Wdc:         false,
	}

	// when
	respDto, err := client.TimetableSearch(reqDto)
	fmt.Printf("Result: %+v\n", respDto)

	// then
	require.Nil(t, err)
	require.NotNil(t, respDto)

	// and response contains at least 1 city
	require.NotEmpty(t, respDto.OutboundFlights)
	assert.NotEmpty(t, respDto.OutboundFlights[0].DepartureStation)
	assert.NotEmpty(t, respDto.OutboundFlights[0].ArrivalStation)
	assert.NotEmpty(t, respDto.OutboundFlights[0].DepartureDate)
	assert.NotEmpty(t, respDto.OutboundFlights[0].DepartureDates)

	require.NotNil(t, respDto.OutboundFlights[0].Price)
	assert.NotEmpty(t, respDto.OutboundFlights[0].Price.Amount)
	assert.NotEmpty(t, respDto.OutboundFlights[0].Price.CurrencyCode)
}
