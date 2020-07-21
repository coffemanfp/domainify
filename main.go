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
	tldsArray := strings.Split(strings.Trim(tlds, " "), ",")

	rand.Seed(time.Now().UTC().UnixNano())
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

func init() {
	flag.StringVar(&tlds, "tlds", "com,net,org,website", "Top-level Domain names")

	flag.Parse()
}
