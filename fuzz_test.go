package main

import (
	"testing"
)

func FuzzParseName(f *testing.F) {
	testcases := []string{"John Doe", "Alice Smith", ""}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		firstName, lastName, err := ParseName(input)
		if err == nil {
			if firstName == "" || lastName == "" {
				t.Errorf("ParseName(%q) = (%q, %q, %v), want non-empty names", input, firstName, lastName, err)
			}
		}
	})
}
