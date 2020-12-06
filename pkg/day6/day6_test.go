package day6

import (
	"fmt"
	"testing"
)

func TestQuestions(t *testing.T) {
	if len(questionID) != 26 {
		t.Errorf("Expecting all 26 letters but got %d", len(questionID))
	}
}

func TestPart1(t *testing.T) {
	println(fmt.Sprintf("Found total with at least one %d", Part1()))
}

func TestPart2(t *testing.T) {
	println(fmt.Sprintf("Found total with all same answer %d", Part2()))
}