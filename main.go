package main

import (
	"errors"
	"strings"
)

func ParseName(input string) (string, string, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return "", "", errors.New("input must contain exactly one space")
	}
	return parts[0], parts[1], nil
}
