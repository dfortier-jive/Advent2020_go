package day14

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"dfortier.org/advent2020/pkg/util"
)

type value struct {
	binaryValue string
}

func newValue(decimalValue int64) *value {
	return &value{
		binaryValue: convertToBinary(decimalValue),
	}
}

func convertToBinary(decimalValue int64) string {
	result := make([]rune, 0)

	for quot := decimalValue; quot > 0; quot = quot / 2 {
		remainder := quot % 2
		if remainder == 1 {
			result = append(result, '1')
		} else {
			result = append(result, '0')
		}
	}

	resultString := ""
	for i := len(result) - 1; i >= 0; i-- {
		resultString += string(result[i])
	}

	return fmt.Sprintf("%036s", resultString)
}

func (v *value) getDecimal() int64 {
	var result int64 = 0
	for i := len(v.binaryValue) - 1; i >= 0; i-- {
		bit := v.binaryValue[i]
		if bit == '1' {
			position := len(v.binaryValue) - 1 - i
			result += int64(math.Pow(2, float64(position)))
		}
	}
	return result
}

func (v *value) applyMask(mask *mask) *value {
	maskedValue := make([]rune, len(v.binaryValue))
	for i := 0; i < len(mask.maskValue); i++ {
		maskValue := mask.maskValue[i]
		switch maskValue {
		case '1':
			maskedValue[i] = '1'
		case '0':
			maskedValue[i] = '0'
		case 'X':
			maskedValue[i] = rune(v.binaryValue[i])
		}
	}
	return &value{
		binaryValue: string(maskedValue),
	}
}

type mask struct {
	maskValue string
}

func newMask(maskString string) *mask {
	return &mask{
		maskValue: maskString,
	}
}

type memory struct {
	memSpot map[int]int64
}

func newMemory() *memory {
	return &memory{
		memSpot: make(map[int]int64, 0),
	}
}

func (m *memory) Apply(address int, value int64) {
	m.memSpot[address] = value
}

func (m *memory) Execute(instructions *instructions) {
	for _, instruction := range instructions.instructions {
		valueBefore := newValue(instruction.value)
		valueAfter := valueBefore.applyMask(instructions.mask)
		m.Apply(instruction.address, valueAfter.getDecimal())
	}
}

func (m *memory) AddAll() int64 {
	var result int64 = 0
	for _, value := range m.memSpot {
		result += value
	}
	return result
}

type instructions struct {
	mask         *mask
	instructions []*instruction
}

type instruction struct {
	address int
	value   int64
}

func newInstruction(address int, value int64) *instruction {
	return &instruction{
		address: address,
		value:   value,
	}
}

func (i *instructions) Add(address int, value int64) {
	i.instructions = append(i.instructions, newInstruction(address, value))
}

func newInstructions(mask *mask) *instructions {
	return &instructions{
		mask:         mask,
		instructions: make([]*instruction, 0),
	}
}

func GetInstructionsFromFile() []*instructions {
	var f *os.File
	var err error
	var result = make([]*instructions, 0)
	if f, err = os.Open("input.txt"); err != nil {
		panic("Unable to read file")
	}

	scanner := bufio.NewScanner(bufio.NewReader(f))
	var currentInstruction *instructions
	for scanner.Scan() {
		line := scanner.Text()

		keyPair := strings.Split(line, " = ")
		if keyPair[0] == "mask" {
			mask := newMask(keyPair[1])
			currentInstruction = newInstructions(mask)
			result = append(result, currentInstruction)
		} else {
			memoryAndAddress := strings.Split(keyPair[0], "[")
			addressString := memoryAndAddress[1]
			addressString = addressString[0 : len(addressString)-1]
			address := util.Convert(addressString)
			value := util.ConvertInt64(keyPair[1])

			currentInstruction.Add(address, value)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}
