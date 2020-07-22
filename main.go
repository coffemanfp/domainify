package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds string

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	tldsArray := splitStringArray(tlds)
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune

		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}

		fmt.Println(string(newText) + "." + tldsArray[rand.Intn(len(tldsArray))])
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

	flag.Parse()
}
