package day5

import (
	"testing"
)

func TestPart1(t *testing.T) {
	Part1()
}

func TestPart2(t *testing.T) {
	Part2()
}

func TestConvertPassport(t *testing.T) {
	testCases := []struct {
		boardingPass      string
		expectedRow int
		expectedCol int
	}{
		{
			boardingPass:      "FBFBBFFRLR",
			expectedRow: 44,
			expectedCol: 5,
		},
		{
			boardingPass: "BFFFBBFRRR",
			expectedRow: 70,
			expectedCol: 7,
		},
	}
	for _, c := range testCases {
		t.Log(c.boardingPass)

		row, col := processPass(c.boardingPass)
		if row != c.expectedRow {
			t.Errorf("Expecting row to be %d, was %d", c.expectedRow, row)
		}
		if col != c.expectedCol {
			t.Errorf("Expecting column to be %d, was %d", c.expectedCol, col)
		}

	}
}

