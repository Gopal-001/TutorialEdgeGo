package main

import (
	"errors"
	"fmt"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func GetMinMax(flights []Flight) (int, int, error) {
	if len(flights) == 0 {
		return 0, 0, errors.New("The list is empty")
	}
	min := flights[0].Price
	max := flights[0].Price
	for id, _ := range flights {
		if flights[id].Price < 0 {
			return 0, 0, errors.New("Price cannot be negative")
		}
		if flights[id].Price < min {
			min = flights[id].Price
		}
		if flights[id].Price > max {
			max = flights[id].Price
		}
	}
	return min, max, nil
}

func main() {
	fmt.Println("Getting the Minimum and Maximum Flight Prices")
	list := []Flight{{Price: 23}, {Price: 54}, {Price: 77}}
	err := []Flight{{Price: 32}, {Price: -5}}
	mn, mx, er := GetMinMax(list)
	fmt.Println(mn, mx, er)
	mn, mx, er = GetMinMax(err)
	fmt.Println(mn, mx, er)
}
