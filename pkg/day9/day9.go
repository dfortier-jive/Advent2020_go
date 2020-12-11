package day9

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"dfortier.org/advent2020/pkg/util"
)

const preambleSize = 25

func readData() []int64 {
	var f *os.File
	var err error
	var result = make([]int64, 0)
	if f, err = os.Open("data.txt"); err != nil {
		panic("Unable to read file")
	}

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()

		value := util.ConvertInt64(line)

		result = append(result, value)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func Part1() (int64, error) {
	data := readData()
	return processMessage(data)
}

func processMessage(data []int64) (int64, error) {
	for i := preambleSize; i < len(data); i++ {
		first := data[i]
		found := false
		println(fmt.Sprintf("Looking for sum of %d (i: %d)", first, i))
		for j := i - 1; j >= (i - preambleSize); j-- {
			for k := j - 1; k >= (i - preambleSize); k-- {
				println(fmt.Sprintf("Summing j,k (%d, %d) = %d + %d", j, k, data[j], data[k]))
				second := data[j] + data[k]
				println(fmt.Sprintf("Sum = %d", second))
				if first == second {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return first, nil
		}
	}
	return -1, errors.New("All numbers are sums of other")
}

func Part2() int64 {
	//invalidValue, _ := Part1()
	invalidValue := int64(2089807806)
	data := readData()

	return findContigentNumbers(invalidValue, data)

}

func findContigentNumbersIndexes(numberToFind int64, data []int64) (int, int, error) {
	for i := 0; i < len(data); i++ {
		sum := data[i]
		for j := i + 1; j < len(data); j++ {
			if sum < numberToFind {
				sum += data[j]
			} else {
				break
			}
			if sum == numberToFind {
				return i, j, nil
			}
		}
	}
	return -1, -1, errors.New("No sum found")
}

func findContigentNumbers(numberToFind int64, data []int64) int64 {
	minIndex, maxIndex, err := findContigentNumbersIndexes(numberToFind, data)
	minValue, maxValue := int64(-1), int64(0)
	for i := minIndex; i <= maxIndex; i++ {
		value := data[i]
		if value > maxValue {
			maxValue = value
		}
		if minValue == -1 || value < minValue {
			minValue = value
		}
	}
	if err != nil {
		panic("No sum found")
	}
	return minValue + maxValue
}
