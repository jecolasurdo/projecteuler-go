package p705_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/jecolasurdo/projecteuler/p705"
	"github.com/stretchr/testify/assert"
)

func Test_InversionCountSumFinal(t *testing.T) {
	t.Skip()
	result := p705.InversionCountSum(10e8)
	assert.Equal(t, 480440153, result)
}

func Test_InversionCountSum(t *testing.T) {
	// testCases := []int{20, 3312, 50, 338079744}
	testCases := []int{20, 3312}

	for i := 0; i < len(testCases); i += 2 {
		n := testCases[i]
		e := testCases[i+1]
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			r := p705.InversionCountSum(n)
			assert.Equal(t, e, r)
		})
	}
}

func Test_InversionCount(t *testing.T) {
	testCases := []int{34214, 5, 21111, 4}

	for i := 0; i < len(testCases); i += 2 {
		n := testCases[i]
		e := testCases[i+1]
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			r := p705.InversionCount([]byte(strconv.Itoa(n)))
			assert.Equal(t, e, r)
		})
	}
}

// func Test_DividedSequence(t *testing.T) {
// 	actual := p705.DividedSequence("432")
// 	expected := []string{
// 		"432", "431", "412", "411",
// 		"232", "231", "212", "211",
// 		"132", "131", "112", "111"}
// 	assert.ElementsMatch(t, expected, actual)
// }

func Test_PrimeConcat(t *testing.T) {
	result := p705.PrimeConcat(20)
	expected := "235711131719"
	assert.Equal(t, expected, result)
}

func Benchmark_InversionCountSum(b *testing.B) {
	scales := []int{5, 10, 20, 30, 40}
	for _, scale := range scales {
		b.Run(fmt.Sprintf("%v", scale), func(b *testing.B) {
			for n := 1; n <= b.N; n++ {
				result := p705.InversionCountSum(scale)
				_ = result
			}
		})
	}
}

func Benchmark_PrimeConcat(b *testing.B) {
	scales := []int{10e1, 10e2, 10e3, 10e4}
	for _, scale := range scales {
		b.Run(fmt.Sprintf("%v", scale), func(b *testing.B) {
			for n := 1; n <= b.N; n++ {
				result := p705.PrimeConcat(scale)
				_ = result
			}
		})
	}
}
