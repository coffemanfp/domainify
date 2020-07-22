package main

import (
	"regexp"
	"strings"
	"unicode"
)

const allowedChars = `[a-z0-9\-]`

// GenNewSecondLevelDomain Generates a new domain name based on the text parameter. Returns a valid RFC 1123 domain name.
func GenNewSecondLevelDomain(text string) (name string) {
	//var invalidCharsPositions map[int]string

	text = strings.ToLower(text)

	re := regexp.MustCompile(`[a-z0-9]`)
	lettersAndNumbersIndexs := re.FindAllStringIndex(text, -1)

	if len(lettersAndNumbersIndexs) == 0 {
		return
	}

	// Delete invalid characters from start
	if lettersAndNumbersIndexs[0][0] != 0 {
		text = text[lettersAndNumbersIndexs[0][0]:]
		lettersAndNumbersIndexs = re.FindAllStringIndex(text, -1)
	}

	// Delete invalid characters until end
	lastLetterOrNumberIndex := lettersAndNumbersIndexs[len(lettersAndNumbersIndexs)-1][0]
	lastTextChar := text[len(text)-1]

	// If the last character is not valid, delete invalid characters until end
	if text[lastLetterOrNumberIndex] != lastTextChar {
		text = text[:lastLetterOrNumberIndex+1]
	}

	var runeName []rune

	// Validate the remaining characters and replace spaces for hyphen
	for _, runeChar := range text {
		if unicode.IsSpace(runeChar) {
			runeChar = '-'
		}

		if runeChar != '-' && !unicode.IsLetter(runeChar) && !unicode.IsNumber(runeChar) {
			continue
		}

		runeName = append(runeName, runeChar)
	}

	name = string(runeName)
	return
}
