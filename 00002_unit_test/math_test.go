package math

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1, 3)
	if sum != 4 {
		t.Errorf("Add(1, 3) failed, expected %v to be 4", sum)
	}
}
