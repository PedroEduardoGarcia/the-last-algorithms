package tests

import (
	"testing"
	"the-last-algorithms/code"
)

func TestAdd(t *testing.T) {
	arr := []any{10, 568, "hello", 45.23, 111}
	if result := code.LinearSearch(&arr, "world"); result != false {
		t.Errorf("LinearSearch(arr, 'world') = %v; want false", result)
	}
	if result := code.LinearSearch(&arr, 11); result != false {
		t.Errorf("LinearSearch(arr,11) = %v; want false", result)
	}
	if result := code.LinearSearch(&arr, "hello"); result != true {
		t.Errorf("LinearSearch(arr, 'hello') = %v; want true", result)
	}
	if result := code.LinearSearch(&arr, 45.23); result != true {
		t.Errorf("LinearSearch(arr, 45.23) = %v; want true", result)
	}
}
