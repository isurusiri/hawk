package main

import (
	"fmt"

	"github.com/isurusiri/hawk/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)

	widescanresults := port.WideScan("localhost")
	fmt.Println(widescanresults)
}
