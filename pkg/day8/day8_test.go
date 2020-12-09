package day8

import (
	"fmt"
	"testing"
)

func TestReadData(t *testing.T) {
	program := readData()

	for _, instruction := range program {
		println(fmt.Sprintf("Doing %s with %d", instruction.instruction, instruction.value))
	}
}

func TestPart1(t *testing.T) {
	loop, value := Part1()

	if !loop {
		println(fmt.Sprintf("Program executed without loop with final value %d", value))
	} else {
		println(fmt.Sprintf("Loop stop with value %d", value))
	}
}
