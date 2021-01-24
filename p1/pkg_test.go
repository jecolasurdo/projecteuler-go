package p1_test

import (
	"strconv"
	"testing"

	"github.com/jecolasurdo/projecteuler/p1"
)

func Test_FindSumOfMultiples(t *testing.T) {
	suts := []func([]int, int) int{
		p1.FindSumOfMultiples,
		p1.FindSumOfMultiplesBruteForce,
	}

	testCases := []struct {
		baseFactors []int
		sumBelow    int
		expResult   int
	}{
		{
			baseFactors: []int{3, 5},
			sumBelow:    10,
			expResult:   23,
		},
		{
			baseFactors: []int{3, 5},
			sumBelow:    1000,
			expResult:   233168,
		},
	}

	for s, sut := range suts {
		for i, testCase := range testCases {
			caseNumber := strconv.Itoa(i)
			t.Run(caseNumber, func(t *testing.T) {
				actResult := sut(testCase.baseFactors, testCase.sumBelow)
				if actResult != testCase.expResult {
					t.Errorf("Test function %v, case %v: Expected %v, got %v", s, caseNumber, testCase.expResult, actResult)
				}
			})
		}
	}

}

func Benchmark_FindSumOfMultiplesBruteForce(b *testing.B) {
	baseFactors := []int{3, 5}
	sumBelow := 1000

	for n := 0; n < b.N; n++ {
		p1.FindSumOfMultiplesBruteForce(baseFactors, sumBelow)
	}
}

func Benchmark_FindSumOfMultiples(b *testing.B) {
	baseFactors := []int{3, 5}
	sumBelow := 1000

	for n := 0; n < b.N; n++ {
		p1.FindSumOfMultiples(baseFactors, sumBelow)
	}
}
