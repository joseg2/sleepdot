//
// sleepdot unix sleep command that prints a dot every second
//
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	const usage string = `Usage: sleepdot {-m STRING} duration (in seconds)`
	var markerPtr string
	flag.StringVar(&markerPtr, "m",".", "Marker to display for each second")
	flag.Parse()

	if flag.NArg() == 1 {
		duration := flag.Arg(0)
		intDuration, err := strconv.Atoi(duration)

		if err != nil {
			fmt.Println("\nERROR: Invalid time provided, duration must be a positive integer")
			fmt.Println(usage)
			flag.PrintDefaults()
			os.Exit(1)
		} else {
			for i:=0 ; i < intDuration ; i++ {
				time.Sleep(time.Second)
				fmt.Printf(markerPtr)
			}
			fmt.Printf("\n")
		}

	} else {
		fmt.Println(usage)
		flag.PrintDefaults()
		os.Exit(1)
	}

}