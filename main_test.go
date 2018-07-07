package main

import (
	"strings"
	"testing"
)

func generateTest(t *testing.T) {
	r := strings.NewReader("test")

	//initializes possible results
	possibleResults := make([]string, len(transforms))
	for i, v := range transforms {
		possibleResults[i] = strings.Replace(v, otherWord, "test", -1)
	}

}
