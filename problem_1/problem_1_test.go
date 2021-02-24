package problem_1

import "testing"

func TestProblem1(t *testing.T) {
	result := TwoOldestAges(t, []int{5, 7, 4, 9, 2})
	t.Logf("Final Result : %v\n", result)
}
