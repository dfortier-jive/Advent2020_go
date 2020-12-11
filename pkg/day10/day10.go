package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"dfortier.org/advent2020/pkg/util"
)

type adapter struct {
	rating int
}

func readData() []adapter {
	var f *os.File
	var err error
	var result = make([]adapter, 0)
	if f, err = os.Open("adapters.txt"); err != nil {
		panic("Unable to read file")
	}

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()

		value := util.Convert(line)
		adapter := adapter{
			rating: value,
		}

		result = append(result, adapter)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func Part1() int {
	fromFile := readData()
	sort.Slice(fromFile,
		func(i, j int) bool {
			return fromFile[i].rating < fromFile[j].rating
		})
	adapters := make([]adapter, 0)
	outlet := adapter{
		rating: 0,
	}
	highestRating := fromFile[len(fromFile)-1].rating + 3
	builtIn := adapter{
		rating: highestRating,
	}
	adapters = append(adapters, outlet)
	adapters = append(adapters, fromFile...)
	adapters = append(adapters, builtIn)
	oneStep, threeStep := 0, 0
	for i := 1; i < len(adapters); i++ {
		if (adapters[i].rating - adapters[i-1].rating) == 1 {
			oneStep++
		}
		if (adapters[i].rating - adapters[i-1].rating) == 3 {
			threeStep++
		}
	}
	return oneStep * threeStep
}
