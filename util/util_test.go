package util

import (
	"testing"
)

func TestIsStringEmpty(t *testing.T) {
	isEmpty := IsStringEmpty("")
	if !isEmpty {
		t.Fatalf(`"" should be %v, %v, found`, true, isEmpty)
	}
}
