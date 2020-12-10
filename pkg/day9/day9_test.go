package day9

import (
	"fmt"
	"testing"
)

func TestReadData(t *testing.T) {
	allNumbers := readData()

	for _, value := range allNumbers {
		println(fmt.Sprintf("Doing %d", value))
	}
}

func TestPart1(t *testing.T) {
	allNumbers := readData()

	if value, err := processMessage(allNumbers); err == nil {
		println(fmt.Sprintf("Found %d that is not sum", value))
		return
	}
	t.Error("Not found")
}
