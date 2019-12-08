package gowizz

import "time"

// GetAllConnections Retrieves all Wizzair connections
func (wizz *WizzClient) GetAllConnections() ([]FlightConnection, error) {
	cities, err := wizz.GetCities()
	if err != nil {
		return nil, err
	}

	var result = make([]FlightConnection, 0, 500)
	for _, departureAirport := range cities.Cities {
		for _, arrivalAirport := range departureAirport.Connections {
			// do something
			result = append(result, FlightConnection{
				DepartureStation: departureAirport.Iata,
				ArrivalStation:   arrivalAirport.Iata,
			})
		}
	}

	return result, nil
}

// GetFlightPrices Get flight prices for a given connection over given number of months
func (wizz *WizzClient) GetFlightPrices(connection FlightConnection, months int) ([]TimetableOutboundFlight, error) {
	var result = make([]TimetableOutboundFlight, 0, 50)
	for _, timeRange := range GenTimeRanges(time.Now(), 30*Day, months) {
		filter := TimetableSearchFilterDto{
			FlightList: []TimetableFlightFilter{
				TimetableFlightFilter{
					DepartureStation: connection.DepartureStation,
					ArrivalStation:   connection.ArrivalStation,
					From:             timeRange.From.Format("2006-01-02"),
					To:               timeRange.To.Format("2006-01-02"),
				},
			},
			AdultCount:  1,
			ChildCount:  0,
			InfantCount: 0,
			Wdc:         false,
		}

		// when
		respDto, err := CopyClient(wizz).TimetableSearch(filter)
		if err != nil {
			return nil, err
		}

		if (len(respDto.OutboundFlights) == 0) {
			return 	result, nil			
		}

		// todo terminate if no more flight
		result = append(result, respDto.OutboundFlights...)
	}

	return result, nil
}
