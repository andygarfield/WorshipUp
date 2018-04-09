package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var songTitleRegex, _ = regexp.Compile(`[a-zA-Z\s\,()]+`)
var songBodyRegex, _ = regexp.Compile(`^[!;. ]([0-9A-Za-z ._,\-–'!?;"’“”#/])*$`)

// ScrubUserTitle takes an input title passes it through a white-list to see if it's valid
func ScrubUserTitle(s string) (string, error) {
	s = strings.Replace(s, "\r\n", "\n", -1)
	if songTitleRegex.Match([]byte(s)) {
		return s, nil
	}
	return "", errors.New("Error: Invalid title")
}

// ScrubUserBody takes an input song body and passes it through a white-list to see if it's valid
func ScrubUserBody(s string) (string, error) {
	s = strings.Replace(s, "\r\n", "\n", -1)
	s = fmt.Sprintf("%s", s)

	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if !songBodyRegex.Match([]byte(line)) && len(line) > 0 {
			return "", fmt.Errorf("Error: Invalid input on line\n%d: %s", i+1, line)
		}
	}
	return s, nil
}
