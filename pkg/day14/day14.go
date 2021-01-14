package day14

import (
	"bufio"
	"container/list"
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

func (v *value) applyAddressMask(mask *mask) *value {
	maskedValue := make([]rune, len(v.binaryValue))
	for i := 0; i < len(mask.maskValue); i++ {
		maskValue := mask.maskValue[i]
		switch maskValue {
		case '1':
			maskedValue[i] = '1'
		case '0':
			maskedValue[i] = rune(v.binaryValue[i])
		case 'X':
			maskedValue[i] = 'X'
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
	memSpot map[int64]int64
}

func newMemory() *memory {
	return &memory{
		memSpot: make(map[int64]int64, 0),
	}
}

func (m *memory) Apply(address int64, value int64) {
	m.memSpot[address] = value
}

func (m *memory) Execute(instructions *instructions) {
	for _, instruction := range instructions.instructions {
		valueBefore := newValue(instruction.value)
		valueAfter := valueBefore.applyMask(instructions.mask)
		m.Apply(int64(instruction.address), valueAfter.getDecimal())
	}
}

func (m *memory) ExecutePart2(instructions *instructions) {
	for _, instruction := range instructions.instructions {
		addressesAfter := getAddresses(instruction.address, instructions.mask)
		for _, addressAfter := range addressesAfter {
			m.Apply(addressAfter.getDecimal(), instruction.value)
		}
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

func getAddresses(initialAddress int, mask *mask) []*value {
	addressValue := newValue(int64(initialAddress))

	valueWithFloating := addressValue.applyAddressMask(mask)
	result := list.New()

	getAddressRec(valueWithFloating.binaryValue, result)

	sliceResult := make([]*value, result.Len())
	i := 0
	for e := result.Front(); e != nil; e = e.Next() {
		sliceResult[i] = e.Value.(*value)
		i++
	}

	return sliceResult
}

func getAddressRec(initialValue string, result *list.List) {
	for i, char := range initialValue {
		switch char {
		case 'X':
			newValueOne := replaceAtIndex(initialValue, '1', i)
			newValueZero := replaceAtIndex(initialValue, '0', i)
			getAddressRec(newValueOne, result)
			getAddressRec(newValueZero, result)
			return
		}
	}
	// No 'X' found, add it
	value := &value{
		binaryValue: initialValue,
	}
	result.PushBack(value)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
