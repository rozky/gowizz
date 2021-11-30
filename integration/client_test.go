package integration

import (
	"fmt"
	"github.com/kr/pretty"
	"github.com/rozky/gowizz"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	debugEnabled = true
)

type ConnectionId struct {
	Departure   string
	Destination string
}

func TestNewCustomClient_shouldCreateNewClient(t *testing.T) {

	// when
	wizz, err := gowizz.NewCustomClient(gowizz.MetadataURL, debugEnabled)

	// then
	if assert.NoError(t, err) {
		assert.NotNil(t, wizz)
	}
}

func TestGetConnections(t *testing.T) {
	cities := getCities()

	connections := cities.GetConnections()

	fmt.Printf("\n\nCount = %d\n\n", len(connections))

	for _, connection := range connections {
		fmt.Printf("%+v\n", connection)
	}
}

func TestGetMultiplePrices(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	cities := getCities()

	//fmt.Printf("Total count = %d\n\n", getConnections(cities))
	//fmt.Printf("Unique count = %d\n\n", len(getConnectionMap(cities)))

	startTime := time.Now()
	count := 0

	for _, departure := range cities.Cities[:10] {
		for _, destination := range departure.Connections {
			outboundId := ConnectionId{
				Departure:   departure.Iata,
				Destination: destination.Iata,
			}

			getPrices(outboundId)
			count++
		}
	}

	//connectionMap := getConnectionMap(cities)
	//for key, _ := range connectionMap {
	//	if processed, ok := connectionMap[key.returnId()]; ok {
	//		getPrices(key, connectionMap[key.returnId()] != nil)
	//	}
	//}

	fmt.Printf("count = %d\ntook %v\n\n", count, time.Now().Sub(startTime))

}

func TestGetCities(t *testing.T) {
	wizz, _ := gowizz.NewCustomClient(gowizz.MetadataURL, debugEnabled)

	// when
	respDto, err := wizz.GetCities()

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
	wizz, _ := gowizz.NewCustomClient(gowizz.MetadataURL, debugEnabled)

	reqDto := gowizz.SearchFilterDto{
		FlightList: []gowizz.FlightFilter{
			{
				DepartureStation: "TAT",
				ArrivalStation:   "LTN",
				DepartureDate:    time.Now().Format("2006-01-02"),
			},
		},
		AdultCount:  1,
		ChildCount:  0,
		InfantCount: 0,
		Wdc:         false,
	}

	// when
	respDto, err := wizz.SearchFlights(reqDto)
	log(respDto)

	// then
	if assert.NoError(t, err) {
		assert.NotNil(t, respDto)
	}
}

func TestTimetableSearch(t *testing.T) {
	wizz, _ := gowizz.NewCustomClient(gowizz.MetadataURL, debugEnabled)

	reqDto := gowizz.TimetableSearchFilterDto{
		FlightList: []gowizz.TimetableFlightFilter{
			{
				DepartureStation: "TAT",
				ArrivalStation:   "LTN",
				From:             time.Now().AddDate(0, 0, 31).Format("2006-01-02"),
				To:               time.Now().AddDate(0, 0, 37).Format("2006-01-02"),
			},
			{
				DepartureStation: "LTN",
				ArrivalStation:   "TAT",
				From:             time.Now().AddDate(0, 0, 31).Format("2006-01-02"),
				To:               time.Now().AddDate(0, 0, 37).Format("2006-01-02"),
			},
		},
		AdultCount:  1,
		ChildCount:  0,
		InfantCount: 0,
		Wdc:         false,
	}

	// when
	respDto, err := wizz.TimetableSearch(reqDto)

	// then
	if assert.NoError(t, err) {
		assert.NotNil(t, respDto)
	}
}

func TestTimetableSearch_MultipleCalls(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	wizz, _ := gowizz.NewCustomClient(gowizz.MetadataURL, debugEnabled)

	reqDto := gowizz.TimetableSearchFilterDto{
		FlightList: []gowizz.TimetableFlightFilter{
			{
				DepartureStation: "TAT",
				ArrivalStation:   "LTN",
				From:             time.Now().AddDate(0, 0, 31).Format("2006-01-02"),
				To:               time.Now().AddDate(0, 0, 61).Format("2006-01-02"),
			},
			{
				DepartureStation: "LTN",
				ArrivalStation:   "TAT",
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
	respDto, err := wizz.TimetableSearch(reqDto)
	assert.NoError(t, err)
	pretty.Println(respDto)

	wizz, _ = gowizz.NewCustomClient(gowizz.MetadataURL, debugEnabled)
	//time.Sleep(time.Duration(5) * time.Second)
	respDto, err = wizz.TimetableSearch(reqDto)
	assert.NoError(t, err)
	pretty.Println(respDto)

	//wizz, _ = NewCustomClient(MetadataURL)
	//respDto, err = wizz.TimetableSearch(reqDto)
	//require.Nil(t, err)
	//log(respDto)

	// wizz, _ = NewCustomClient(MetadataURL)
	// respDto, err = wizz.TimetableSearch(reqDto)
	// require.Nil(t, err)

	// wizz, _ = NewCustomClient(MetadataURL)
	// respDto, err = wizz.TimetableSearch(reqDto)
	// require.Nil(t, err)

	//log(respDto)
	//
	//// then
	//require.Nil(t, err)
	//require.NotNil(t, respDto)
	//
	//// and response contains at least 1 city
	//require.NotEmpty(t, respDto.OutboundFlights)
	//assert.NotEmpty(t, respDto.OutboundFlights[0].DepartureStation)
	//assert.NotEmpty(t, respDto.OutboundFlights[0].ArrivalStation)
	//assert.NotEmpty(t, respDto.OutboundFlights[0].DepartureDate)
	//assert.NotEmpty(t, respDto.OutboundFlights[0].DepartureDates)
	//
	//require.NotNil(t, respDto.OutboundFlights[0].Price)
	//assert.NotEmpty(t, respDto.OutboundFlights[0].Price.Amount)
	//assert.NotEmpty(t, respDto.OutboundFlights[0].Price.CurrencyCode)
}

func getPrices(outboundId ConnectionId) {
	wizz, _ := gowizz.NewCustomClient(gowizz.MetadataURL, debugEnabled)

	reqDto := gowizz.TimetableSearchFilterDto{
		FlightList: []gowizz.TimetableFlightFilter{
			{
				DepartureStation: outboundId.Departure,
				ArrivalStation:   outboundId.Destination,
				From:             time.Now().AddDate(0, 0, 31).Format("2006-01-02"),
				To:               time.Now().AddDate(0, 0, 61).Format("2006-01-02"),
			},
			{
				DepartureStation: outboundId.Destination,
				ArrivalStation:   outboundId.Departure,
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
	if respDto, err := wizz.TimetableSearch(reqDto); err != nil {
		panic(err)
	} else {
		log(respDto)
	}
}

func getCities() *gowizz.CitiesDto {
	wizz, _ := gowizz.NewCustomClient(gowizz.MetadataURL, debugEnabled)

	// when
	if cities, err := wizz.GetCities(); err != nil {
		panic(err)
	} else {
		return cities
	}
}

func log(resp interface{}) {
	if debugEnabled {
		pretty.Println(resp)
	}
}
