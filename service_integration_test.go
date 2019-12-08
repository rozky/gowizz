package gowizz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllConnections(t *testing.T) {
	client, _ := NewClient(MetadataURL)

	// when
	connections, err := client.GetAllConnections()
	fmt.Printf("%d\n", len(connections))
	fmt.Printf("%+v\n", connections)

	// then
	require.Nil(t, err)
	require.NotNil(t, connections)

}

func TestGetFlightPrices(t *testing.T) {
	client, _ := NewClient(MetadataURL)
	connection := FlightConnection{
		DepartureStation: "LTN",
		ArrivalStation:   "BTS",
		// ArrivalStation:   "TAT",
	}

	// when
	connections, err := client.GetFlightPrices(connection, 15)
	fmt.Printf("%d\n", len(connections))
	fmt.Printf("%+v\n", connections)

	// then
	require.Nil(t, err)
	require.NotNil(t, connections)
}
