package projecteuler_test

import (
	"testing"

	"github.com/jecolasurdo/projecteuler"
	"github.com/stretchr/testify/assert"
)

func Test_InversionCountSum(t *testing.T) {
	result := projecteuler.InversionCountSum(20)
	expected := 3312
	assert.Equal(t, expected, result)

	result = projecteuler.InversionCountSum(50)
	expected = 338079744
	assert.Equal(t, expected, result)
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
