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
