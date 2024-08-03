package tests

import (
	"testing"
	"the-last-algorithms/code"
)

func TestAdd(t *testing.T) {
	if result := code.Add(1, 2); result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
	if result := code.Add(-1, 1); result != 0 {
		t.Errorf("Add(-1, 1) = %d; want 0", result)
	}
	if result := code.Add(-1, -1); result != -2 {
		t.Errorf("Add(-1, -1) = %d; want -2", result)
	}
}
