package day15

import "fmt"

type number struct {
	number       int
	lastTime     int
	lastLastTime int
}

func newNumber(aNumber int) *number {
	return &number{
		number:       aNumber,
		lastTime:     -1,
		lastLastTime: -1,
	}
}

func sayNumbers(values []int) int {
	numbers := make(map[int]*number)

	countSpoken := 1
	lastNumber := -1
	for _, aValue := range values {
		numbers[aValue] = newNumber(aValue)

		numbers[aValue].lastTime = countSpoken
		println(fmt.Sprintf("Turn %d: count %d", countSpoken, aValue))
		countSpoken++
		lastNumber = aValue
	}

	starterNumber := 0
	for ; countSpoken <= 2020; countSpoken++ {
		if number, ok := numbers[lastNumber]; ok {
			if number.lastLastTime == -1 {
				lastNumber = values[starterNumber]
			} else {
				lastNumber = number.lastTime - number.lastLastTime
			}
			println(fmt.Sprintf("Turn %d: count %d", countSpoken, lastNumber))

			if _, ok := numbers[lastNumber]; !ok {
				numbers[lastNumber] = newNumber(lastNumber)
			}
			// Note where we are at
			numbers[lastNumber].lastLastTime = numbers[lastNumber].lastTime
			numbers[lastNumber].lastTime = countSpoken
		} else {
			// Say a started number
			lastNumber = values[starterNumber]

			//starterNumber = (starterNumber + 1) % len(values)
			println(fmt.Sprintf("Turn %d: count %d", countSpoken, lastNumber))

			// Note where we are at
			numbers[lastNumber].lastLastTime = numbers[lastNumber].lastTime
			numbers[lastNumber].lastTime = countSpoken
		}
	}
	return lastNumber
}
