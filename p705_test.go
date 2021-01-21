package projecteuler_test

import (
	"testing"

	"github.com/jecolasurdo/projecteuler"
	"github.com/stretchr/testify/assert"
)

func Test_InversionCount(t *testing.T) {
	result := projecteuler.InversionCount("34214")
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func Test_DividedSequence(t *testing.T) {
	result := projecteuler.DividedSequence("332")
	expected := []string{"332", "331", "312", "311", "132", "131", "112", "111"}
	assert.ElementsMatch(t, expected, result)
}

func Test_PrimeConcat(t *testing.T) {
	result := projecteuler.PrimeConcat(20)
	expected := "235711131719"
	assert.Equal(t, expected, result)
}
