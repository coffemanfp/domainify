package main

import "regexp"

// ValidateDomainName Validates a domain name.
func ValidateDomainName(name string) (valid bool) {
	if name == "" || len(name) > 253 {
		return
	}

	rx := regexp.MustCompile(`^ ([a-zA-Z0-9] | [a-zA-Z0-9] [a-zA-Z0-9 \ -] {0,61} [a-zA-Z0-9]) (\. ([a-zA-Z0-9] | [a-zA-Z0-9] [a-zA-Z0-9 \ -] {0,61} [a-zA-Z0-9])) * $`)

	valid = rx.MatchString(name)
	return
}
