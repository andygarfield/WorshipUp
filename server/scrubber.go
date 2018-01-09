package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var songBodyRegex, _ = regexp.Compile(`^(!|;|\.|\s)[a-zA-Z\s.\,\;]+`)
var songTitleRegex, _ = regexp.Compile(`[a-zA-Z\s\,()]+`)

func scrubUserData(s string) (string, error) {
	s = strings.Replace(s, "\r", "\n", -1)
	if songBodyRegex.Match([]byte(s)) {
		fmt.Println(s)
		return s, nil
	}
	return "", errors.New("Invalid input")
}

func scrubUserTitle(s string) (string, error) {
	s = strings.Replace(s, "\r", "\n", -1)
	if songTitleRegex.Match([]byte(s)) {
		fmt.Println(s)
		return s, nil
	}
	return "", errors.New("Invalid input")
}
