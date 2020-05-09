package gowizz

import (
	"math/rand"
)

const (
	// WizzBaseURL The URL of live wizzair website
	WizzBaseURL = "https://wizzair.com/"

	// MetadataURL MetadataURL
	MetadataURL = WizzBaseURL + "/static_fe/metadata.json"

	// CitiesPath Endpoint that returns information about cities Wizzair flight to/from
	CitiesPath = "/asset/map?languageCode=en-gb&forceJavascriptOutput=false"

	// SearchPath Search endpoint
	SearchPath = "/search/search"

	// TimetableSearchPath Timetable search endpoint
	TimetableSearchPath = "/search/timetable"

	// UserAgent User Agent to use when calling Wizzair (without it calls take about 10s but with it only 100ms)
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"
)

// MetadataDto HTTP body of GET <MetadataURL> call
type MetadataDto struct {
	ApiURL string `json:"apiUrl"`
}

// CitiesDto HTTP body of GET <CitiesPath> call
type CitiesDto struct {
	Cities []City `json:"cities"`
}

func (cities *CitiesDto) FindCity(iota string) (City, bool) {
	for _, city := range cities.Cities {
		if city.Iata == iota {
			return city, true
		}
	}

	return City{}, false
}

func (cities *CitiesDto) ConnectionExists(fromIota string, toIota string) bool {
	if departure, ok := cities.FindCity(fromIota); ok {
		return departure.IsConnectedTo(toIota)
	}

	return false
}

func (cities *CitiesDto) GetConnections() []Connection {
	var result []Connection

	processed := map[string]bool{}
	for _, departure := range cities.Cities {
		processed[departure.Iata] = true

		for _, destination := range departure.Connections {
			if _, contains := processed[destination.Iata]; !contains {
				bothWays := cities.ConnectionExists(destination.Iata, departure.Iata)
				result = append(result, Connection{Departure: departure.Iata, Destination: destination.Iata, BothWays: bothWays})
			}
		}
	}

	return result
}

// City Contains basic information about city Wizzair flight to/from
type City struct {
	Iata        string          `json:"iata"`
	ShortName   string          `json:"shortName"`
	CountryName string          `json:"countryName"`
	CountryCode string          `json:"countryCode"`
	Longitude   float32         `json:"longitude"`
	Latitude    float32         `json:"latitude"`
	Connections []ConnectedCity `json:"connections"`
}

func (city City) IsConnectedTo(iota string) bool {
	for _, connection := range city.Connections {
		if connection.Iata == iota {
			return true
		}
	}

	return false
}

// ConnectedCity Contains informations about connected city
type ConnectedCity struct {
	Iata string `json:"iata"`
}

// SearchFilterDto Search request body
type SearchFilterDto struct {
	FlightList  []FlightFilter `json:"flightList"`
	AdultCount  int            `json:"adultCount"`
	ChildCount  int            `json:"childCount"`
	InfantCount int            `json:"infantCount"`
	Wdc         bool           `json:"wdc"`
}

// FlightFilter Flight filter in the search request
type FlightFilter struct {
	DepartureStation string `json:"departureStation"`
	ArrivalStation   string `json:"arrivalStation"`
	DepartureDate    string `json:"departureDate"`
}

// SearchResultDto Result of the flight search
type SearchResultDto struct {
	OutboundFlights []OutboundFlight `json:"outboundFlights"`
}

type OutboundFlight struct {
	DepartureStation string       `json:"departureStation"`
	ArrivalStation   string       `json:"arrivalStation"`
	DepartureDate    string       `json:"departureDate"`
	ArrivalDateTime  string       `json:"arrivalDateTime"`
	Duration         string       `json:"duration"`
	Fares            []FlightFare `json:"fares"`
}

type FlightFare struct {
	FareSellKey      string      `json:"fareSellKey"`
	DepartureStation string      `json:"departureStation"`
	AvailableCount   int         `json:"availableCount"`
	SoldOut          bool        `json:"soldOut"`
	BasePrice        FlightPrice `json:"basePrice"`
}

type FlightPrice struct {
	Amount       float32 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
}

type TimetableSearchFilterDto struct {
	FlightList  []TimetableFlightFilter `json:"flightList"`
	AdultCount  int                     `json:"adultCount"`
	ChildCount  int                     `json:"childCount"`
	InfantCount int                     `json:"infantCount"`
	Wdc         bool                    `json:"wdc"`
}

type TimetableFlightFilter struct {
	DepartureStation string `json:"departureStation"`
	ArrivalStation   string `json:"arrivalStation"`
	From             string `json:"from"`
	To               string `json:"to"`
}

type TimetableSearchResultDto struct {
	OutboundFlights []TimetableOutboundFlight `json:"outboundFlights"`
	ReturnFlights   []TimetableOutboundFlight `json:"returnFlights"`
}

type TimetableOutboundFlight struct {
	DepartureStation string      `json:"departureStation"`
	ArrivalStation   string      `json:"arrivalStation"`
	DepartureDate    string      `json:"departureDate"`
	DepartureDates   []string    `json:"departureDates"`
	Price            FlightPrice `json:"price"`
}

type Connection struct {
	Departure   string
	Destination string
	BothWays    bool
}

func GetRandomUserAgent() string {
	agents := []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36",
		"Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 7.0; SM-G892A Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/60.0.3112.107 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 7.0; SM-G930VC Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/58.0.3029.83 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0.1; SM-G935S Build/MMB29K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0.1; SM-G920V Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 5.1.1; SM-G928X Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.83 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 6P Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.83 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0.1; E6653 Build/32.2.A.0.253) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0; HTC One X10 Build/MRA58K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/61.0.3163.98 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0; HTC One M9 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.3",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/69.0.3497.105 Mobile/15E148 Safari/605.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/13.2b11866 Mobile/16A366 Safari/605.1.15",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A5370a Safari/604.1",
		"Mozilla/5.0 (iPhone9,3; U; CPU iPhone OS 10_0_1 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/14A403 Safari/602.1",
		"Mozilla/5.0 (iPhone9,4; U; CPU iPhone OS 10_0_1 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/14A403 Safari/602.1",
	}

	return agents[rand.Intn(len(agents))]
}
