package day6

import (
	"fmt"
	"testing"
)

func TestReadData(t *testing.T) {
	data := readData()
	if 5 != len(data) {
		t.Errorf("Expecting %d but was %d", 5, len(data))
	}

	for _, oneForm := range data {
		println(oneForm.getNbYes())
	}
}

func TestPart1(t *testing.T) {
	println(fmt.Sprintf("Found total %d", Part1()))
}