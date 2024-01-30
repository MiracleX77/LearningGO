package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expectResult := 5
	if result != expectResult {
		t.Errorf("Add(2,3) = %d is wrong, coorect is %d", result, expectResult)
	}
}
