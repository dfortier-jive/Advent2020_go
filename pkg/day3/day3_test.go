package day3


import (
	"fmt"
	"testing"
)

func TestReadData(t *testing.T) {
	result := readData();

	if len(result) != 323 {
		t.Error("Not 323 long")
	}
	if len(result[0]) != 31 {
		t.Error("Not 31 large")
	}
	expected := []bool {false, true, true, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, true, false, false, false, true, false,}
	for i, elem := range result[0] {
		if expected[i] != result[0][i] {
			t.Errorf("Position %d Expecting %t but got %t", i, elem, result[i][0])
		}
	}
	
}

func TestIterateMazePart1(t *testing.T) {
	println(IterateMaze(3, 1))
}


func TestIterateMazePart2(t *testing.T) {
	first := IterateMaze(1, 1)
	second := IterateMaze(3, 1)
	third := IterateMaze(5, 1)
	fourth := IterateMaze(7, 1)
	fifth := IterateMaze(1, 2)

	println(fmt.Sprintf("Found %d", first * second * third * fourth * fifth))
}