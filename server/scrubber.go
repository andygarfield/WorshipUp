package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var songTitleRegex, _ = regexp.Compile(`[a-zA-Z\s\,()]+`)

// Would remove "%3B" if https://github.com/golang/go/issues/23447 is addressed
var songBodyRegex, _ = regexp.Compile(`^[!;. ]([0-9A-Za-z .,\-'!?;"’/]|%3B)*$`)

func scrubUserTitle(s string) (string, error) {
	s = strings.Replace(s, "\r", "\n", -1)
	if songTitleRegex.Match([]byte(s)) {
		return s, nil
	}
	return "", errors.New("Error: Invalid title")
}

func scrubUserData(s string) (string, error) {
	s = strings.Replace(s, "\r", "\n", -1)
	s = fmt.Sprintf("%s", s)

	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if !songBodyRegex.Match([]byte(line)) && len(line) > 0 {
			return "", fmt.Errorf("Error: Invalid input on line\n%d: %s", i+1, line)
		}
	}
	return s, nil
}
