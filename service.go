package gowizz

import (
	"sort"
	"time"

	"sync"
)

// ByDepartureDate implements sort.Interface based on the ByDepartureDate field
type ByDepartureDate []TimetableOutboundFlight

func (a ByDepartureDate) Len() int           { return len(a) }
func (a ByDepartureDate) Less(i, j int) bool { return a[i].DepartureDate < a[j].DepartureDate }
func (a ByDepartureDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// GetAllConnections Retrieves all Wizzair connections
func (wizz *WizzClient) GetAllConnections() ([]FlightConnection, error) {
	cities, err := wizz.GetCities()
	if err != nil {
		return nil, err
	}

	var connections = make([]FlightConnection, 0, 500)
	for _, departureAirport := range cities.Cities {
		for _, arrivalAirport := range departureAirport.Connections {
			// do something
			connections = append(connections, FlightConnection{
				DepartureStation: departureAirport.Iata,
				ArrivalStation:   arrivalAirport.Iata,
			})
		}
	}

	return connections, nil
}

// GetPricesParallel Get flight prices for a given connection over given number of months
func (wizz *WizzClient) GetPricesParallel(flight FlightConnection, months int) ([]TimetableOutboundFlight, error) {

	outChannel := make(chan []TimetableOutboundFlight)

	var wg sync.WaitGroup

	for _, tRange := range GenTimeRanges(time.Now(), 30*Day, months) {
		wg.Add(1)
		go wizz.GetPrices(flight, tRange, outChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(outChannel)
	}()

	var result = make([]TimetableOutboundFlight, 0, 50)
	for prices := range outChannel {
		result = append(result, prices...)
	}

	sort.Sort(ByDepartureDate(result))

	return result, nil
}

// GetPrices Get a flight prices for single time range
func (wizz *WizzClient) GetPrices(flight FlightConnection, tRange TimeRange, out chan []TimetableOutboundFlight, wg *sync.WaitGroup) {
	defer wg.Done()

	filter := TimetableSearchFilterDto{
		FlightList: []TimetableFlightFilter{
			TimetableFlightFilter{
				DepartureStation: flight.DepartureStation,
				ArrivalStation:   flight.ArrivalStation,
				From:             tRange.From.Format("2006-01-02"),
				To:               tRange.To.Format("2006-01-02"),
			},
		},
		AdultCount:  1,
		ChildCount:  0,
		InfantCount: 0,
		Wdc:         false,
	}

	prices, err := wizz.TimetableSearch(filter)
	if err != nil {
		panic(err)
	}

	out <- prices.OutboundFlights
}
