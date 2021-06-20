package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/isurusiri/hawk/port"
)

var (
	tcp     = flag.Bool("tcp", false, "scannes ports listening via tcp")
	udp     = flag.Bool("udp", false, "scannes ports listening via udp")
	all     = flag.Bool("all", true, "scannes all ports")
	help    = flag.Bool("help", false, "display the help options")
	svcmode = flag.Bool("svcmode", false, "specifies to enable service mode")
)

func main() {

	var hostname  string
	var rangeVal  string
	flag.StringVar(&rangeVal, "range", "", "port range to be scanned in the format of 00-00000")
	flag.StringVar(&hostname, "host", "", "host to be scanned")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *tcp {
		fmt.Println("TCP")
	}

	if &rangeVal != nil {
		fmt.Println(rangeVal)
	}

	if &hostname != nil {
		fmt.Println(hostname)
	}

	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)

	widescanresults := port.WideScan("localhost")
	fmt.Println(widescanresults)
}
