package day14

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinaryConversion(t *testing.T) {
	buf := make([]byte, 36)

	for _, x := range []int64{0, 1, 2, 63, 64} {
		n := binary.PutVarint(buf, x)
		fmt.Printf("Got %d bytes", n)
		fmt.Printf("%x\n", buf)
		fmt.Printf("%b\n", buf)
	}
}

func TestDecimalConversion(t *testing.T) {
	testCases := []struct {
		binary   string
		intValue int64
	}{
		{"000000000000000000000000000000000000", 0},
		{"000000000000000000000000000000000001", 1},
		{"000000000000000000000000000000011100", 28},
		{"111111111111111111111111111111111111", 68719476735},
	}

	for _, c := range testCases {
		value := &value{
			binaryValue: c.binary,
		}
		converted := value.getDecimal()
		if converted != c.intValue {
			t.Errorf("Expected %d but got %d", c.intValue, converted)
		}

		convertedFromBin := convertToBinary(c.intValue)
		if convertedFromBin != c.binary {
			t.Errorf("Expected %s but got %s", c.binary, convertedFromBin)
		}
	}
}

func TestMask(t *testing.T) {
	testCases := []struct {
		mask                 string
		intValue             int64
		expectedBinaryString string
	}{
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 11, "000000000000000000000000000001001001"},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 101, "000000000000000000000000000001100101"},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 0, "000000000000000000000000000001000000"},
	}

	for _, c := range testCases {
		value := newValue(c.intValue)
		mask := newMask(c.mask)

		value.applyMask(mask)

		if value.binaryValue != c.expectedBinaryString {
			t.Errorf("Expected %s but got %s", c.expectedBinaryString, value.binaryValue)
		}
	}
}

func TestPart1Example(t *testing.T) {
	memory := newMemory()

	maskString := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	maskObj := newMask(maskString)
	instructions := newInstructions(maskObj)
	instructions.Add(8, 11)
	instructions.Add(7, 101)
	instructions.Add(8, 0)
	memory.Execute(instructions)

	result := memory.AddAll()

	if result != 165 {
		t.Errorf("Expected %d but got %d", 165, result)
	}
}

func TestPart1(t *testing.T) {
	memory := newMemory()
	instructionsSet := GetInstructionsFromFile()

	for _, instructions := range instructionsSet {
		memory.Execute(instructions)
	}

	result := memory.AddAll()

	fmt.Printf("Got sum of %d \n", result)
}
