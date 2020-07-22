package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var tlds string
var hostnames string

func main() {
	var randomTLD string
	var randomHostname string

	tldsArray := splitStringArray(tlds)
	hostnamesArray := splitStringArray(hostnames)
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		sld := GenNewSecondLevelDomain(s.Text())
		randomTLD = tldsArray[rand.Intn(len(tldsArray))]
		randomHostname = hostnamesArray[rand.Intn(len(hostnamesArray))]
		result := randomHostname + "." + sld + "." + randomTLD

		if len(result) > 255 {
			quantityToCut := 255 - len(randomHostname) - len(randomTLD) - 2
			cuttedNewText := sld[:quantityToCut]
			result = randomHostname + "." + cuttedNewText + "." + randomTLD
		}

		// Add random top level domain name and print it at standard output.
		fmt.Println(result)
	}
}

func splitStringArray(stringArray string) (array []string) {
	array = strings.Split(strings.Replace(stringArray, " ", "", -1), ",")
	return
}

func init() {
	// New random generator
	rand.Seed(time.Now().UTC().UnixNano())

	// Flags config
	flag.StringVar(&tlds, "tlds", "com,net,org,website", "Top-level Domain names")
	flag.StringVar(&hostnames, "h", "www,es,en", "hostname")

	flag.Parse()
}
