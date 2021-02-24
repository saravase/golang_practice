package problem_1

import (
	"testing"
)

func TwoOldestAges(t *testing.T, ages []int) [2]int {
	var result [2]int
	secondOldestAge := 0
	oldestAge := secondOldestAge
	for age := range ages {
		if age > oldestAge {
			secondOldestAge = oldestAge
			oldestAge = age
		} else if age > secondOldestAge {
			secondOldestAge = age
		}
	}

	result[0] = secondOldestAge
	result[1] = oldestAge

	return result
}
