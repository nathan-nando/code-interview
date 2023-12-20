package test

import (
	"github.com/bloomberg/go-testgroup"
	"testing"
)

func topFrequentElementsLogic(nums []int, k int) []int {
	var result []int
	temp := make(map[int]int)
	first := nums[0]

	for _, num := range nums {
		if temp[num] == 0 {
			temp[num] = 1
		} else {
			temp[num] = temp[num] + 1
		}
	}

	for i := 0; i < k; i++ {
		result = append(result, findHighestValue(temp, &first))
		delete(temp, first)
	}

	return result
}

func findHighestValue(maps map[int]int, current *int) int {
	for key := range maps {
		if maps[key] > maps[*current] {
			*current = key
		}
	}
	return *current
}

type TopFrequentElements struct {
}

func TestTopFrequentElements(t *testing.T) {
	testgroup.RunSerially(t, &TopFrequentElements{})
}

func (r *TopFrequentElements) TestCase1(test *testgroup.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}

	test.Equal(expected, topFrequentElementsLogic(nums, k))
}

func (r *TopFrequentElements) TestCase2(test *testgroup.T) {
	nums := []int{1}
	k := 1
	expected := []int{1}

	test.Equal(expected, topFrequentElementsLogic(nums, k))
}
