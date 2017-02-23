package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Print("Meters:\t", meters, "\nyards:\t", yards, "\n")
}
