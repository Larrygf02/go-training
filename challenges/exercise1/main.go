package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type byPrice []Flight

func (f byPrice) Len() int {
	return len(f)
}

func (f byPrice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f byPrice) Less(i, j int) bool {
	return f[i].Price < f[j].Price
}

// SortByPrice sorts flights from lower to highest
func SortByPrice(flights []Flight) []Flight {
	// implement
	sort.Sort(byPrice(flights))
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
		fmt.Println("==")
	}
}

func main() {
	flights := []Flight{
		Flight{Origin: "a", Destination: "b", Price: 15},
		Flight{Origin: "c", Destination: "d", Price: 12},
		Flight{Origin: "e", Destination: "f", Price: 20},
		Flight{Origin: "g", Destination: "h", Price: 40},
		Flight{Origin: "a", Destination: "f", Price: 17},
	}
	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
