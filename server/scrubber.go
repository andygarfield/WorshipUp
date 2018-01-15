package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var songTitleRegex, _ = regexp.Compile(`[a-zA-Z\s\,()]+`)
var songBodyRegex, _ = regexp.Compile(`^[!;. \n][A-Za-z .,;\-'\n]+$`)

func scrubUserTitle(s string) (string, error) {
	s = strings.Replace(s, "\r", "\n", -1)
	if songTitleRegex.Match([]byte(s)) {
		return s, nil
	}
	return "", errors.New("Invalid input")
}

func scrubUserData(s string) (string, error) {
	s = strings.Replace(s, "\r", "\n", -1)

	fmt.Println(s)
	lines := strings.Split(s, "\n")
	s = fmt.Sprintf("%s", s)
	fmt.Println(s)

	for _, line := range lines {
		if !songBodyRegex.Match([]byte(line)) {
			for range line {
				return "", errors.New("Invalid input")
			}
		}
	}

	return s, nil
}
