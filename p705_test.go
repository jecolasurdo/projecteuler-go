package projecteuler_test

import (
	"fmt"
	"testing"

	"github.com/jecolasurdo/projecteuler"
	"github.com/stretchr/testify/assert"
)

func Test_InversionCountSum(t *testing.T) {
	testCases := []int{20, 3312, 50, 338079744}

	for i := 0; i < len(testCases); i += 2 {
		n := testCases[i]
		e := testCases[i+1]
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			r := projecteuler.InversionCountSum(n)
			assert.Equal(t, e, r)
		})
	}
}

func Test_InversionCount(t *testing.T) {
	result := projecteuler.InversionCount("34214")
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}

	result = projecteuler.InversionCount("21111")
	if result != 4 {
		t.Errorf("Expected 4, got %v", result)
	}
}

func Test_DividedSequence(t *testing.T) {
	result := projecteuler.DividedSequence("432")
	expected := []string{
		"432", "431", "412", "411",
		"232", "231", "212", "211",
		"132", "131", "112", "111"}
	assert.ElementsMatch(t, expected, result)
}

func Test_PrimeConcat(t *testing.T) {
	result := projecteuler.PrimeConcat(20)
	expected := "235711131719"
	assert.Equal(t, expected, result)
}
