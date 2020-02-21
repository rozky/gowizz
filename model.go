package gowizz

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
