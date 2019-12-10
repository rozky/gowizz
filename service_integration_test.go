package gowizz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllConnections(t *testing.T) {
	wizz, _ := NewCustomClient(MetadataURL)

	// when
	connections, err := wizz.GetAllConnections()
	fmt.Printf("%d\n", len(connections))
	fmt.Printf("%+v\n", connections)

	// then
	require.Nil(t, err)
	require.NotNil(t, connections)

}

func TestGetFlightPrices(t *testing.T) {
	wizz, _ := NewCustomClient(MetadataURL)
	connection := FlightConnection{
		DepartureStation: "LTN",
		ArrivalStation:   "BTS",
		// ArrivalStation:   "TAT",
	}

	// when
	connections, err := wizz.GetFlightPrices(connection, 20)
	fmt.Printf("%d\n", len(connections))
	fmt.Printf("%+v\n", connections)

	// then
	require.Nil(t, err)
	require.NotNil(t, connections)
}
