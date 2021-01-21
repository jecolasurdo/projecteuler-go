package projecteuler_test

import (
	"testing"

	"github.com/jecolasurdo/projecteuler"
)

func Test_InversionCount(t *testing.T) {
	result := projecteuler.InversionCount("34214")
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}
