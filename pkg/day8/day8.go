package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dfortier.org/advent2020/pkg/util"
)

var validInstructions = []string{"nop", "jmp", "acc"}

type instruction struct {
	instruction string
	value       int
	executed    bool
}

func readData() []*instruction {
	var f *os.File
	var err error
	var instructionObj *instruction
	var result = make([]*instruction, 0)
	if f, err = os.Open("instructions.txt"); err != nil {
		panic("Unable to read file")
	}

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		instructionAndValue := strings.Split(line, " ")
		instructionString := instructionAndValue[0]
		valueString := instructionAndValue[1]
		instructionObj = &instruction{
			instruction: instructionString,
			value:       util.Convert(valueString),
			executed:    false,
		}
		result = append(result, instructionObj)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func Part1() (bool, int) {
	program := readData()

	executionPointer := 0
	value := 0
	for true {
		oneInstruction := program[executionPointer]
		if oneInstruction.executed {
			return true, value
		}
		oneInstruction.executed = true

		switch oneInstruction.instruction {
		case "nop":
			executionPointer++
		case "acc":
			executionPointer++
			value += oneInstruction.value
		case "jmp":
			executionPointer += oneInstruction.value
			if executionPointer < 0 || executionPointer >= len(program) {
				panic("Jmp too far")
			}
		}
		if executionPointer == len(program)-1 {
			// Program completed last instruction
			break
		}
	}
	return false, value
}
