package day15

import "testing"

func TestSpokeNumberExample1(t *testing.T) {
	lastNumber := sayNumbers([]int{0, 3, 6})

	if lastNumber != 436 {
		t.Errorf("Expected %d but was %d", 436, lastNumber)
	}
}

func TestSpokeNumberExample2(t *testing.T) {
	lastNumber := sayNumbers([]int{1, 3, 2})

	if lastNumber != 1 {
		t.Errorf("Expected %d but was %d", 1, lastNumber)
	}
}

func TestSpokeNumberExample3(t *testing.T) {
	lastNumber := sayNumbers([]int{2, 1, 3})

	if lastNumber != 10 {
		t.Errorf("Expected %d but was %d", 10, lastNumber)
	}
}
