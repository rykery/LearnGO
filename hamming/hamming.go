// Package hamming allows calculating the Hamming Distance
package hamming

import (
	"errors"
)

// Distance calculates the Hamming Distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("cannot compare distance with unequal lengths")
	}
	cnt := 0
	for i := range a {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt, nil
}
