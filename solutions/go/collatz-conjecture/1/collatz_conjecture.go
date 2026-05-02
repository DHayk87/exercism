package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	currentValue := n
	steps := 0
	if n <= 0 {
		return 0, fmt.Errorf("Error")
	}

	for currentValue != 1 {
		if currentValue%2 == 0 {
			currentValue = currentValue / 2
		} else {
			currentValue = currentValue*3 + 1
		}
		steps++
	}

	return steps, nil
}
