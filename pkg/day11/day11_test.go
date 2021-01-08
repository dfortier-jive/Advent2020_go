package day11

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {

	occupiedSeats := Part1()

	println(fmt.Sprintf("Found %d seats", occupiedSeats))
}
