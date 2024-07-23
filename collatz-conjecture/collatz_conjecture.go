package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n < 1 {
		return 0, errors.New("invalid input")
	}
	for n > 1 {
		if n%2 == 0 {
			steps++
			n = n / 2
		} else {
			steps++
			n = 3*n + 1
		}
	}
	return steps, nil
}
