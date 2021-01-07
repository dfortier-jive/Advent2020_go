package day11

import (
	"bufio"
	"fmt"
	"os"
)

const floor = 0
const emptySeat = 1
const occupied = 2

const seatingHeight = 10

var floorSpot = newSpot(floor)

type spot struct {
	state int
}

func newSpot(state int) *spot {
	return &spot{
		state: state,
	}
}

func dupSpot(oneSpot *spot) *spot {
	return &spot{
		state: oneSpot.state,
	}
}

func (s *spot) isOccupied() bool {
	return s.state == occupied
}

func (s *spot) isFloor() bool {
	return s.state == floor
}

func readData() [][]*spot {
	var f *os.File
	var err error
	if f, err = os.Open("seatings.txt"); err != nil {
		panic("Unable to read file")
	}
	result := make([][]*spot, seatingHeight)

	scanner := bufio.NewScanner(bufio.NewReader(f))
	countLine := 0
	for scanner.Scan() {
		line := scanner.Text()
		result[countLine] = make([]*spot, len(line))

		for i, value := range line {
			switch value {
			case 'L':
				result[countLine][i] = newSpot(emptySeat)
			case '#':
				result[countLine][i] = newSpot(occupied)
			case '.':
				result[countLine][i] = newSpot(floor)
			}
		}

		countLine++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func Part1() {
	seatingSpots := readData()
	printMe(seatingSpots)
	newSeatingSpots := copySeatings(seatingSpots)

	maxIteration := 10
	count := 0

	for iterateSeating(seatingSpots, &newSeatingSpots) && count < maxIteration {
		printMe(newSeatingSpots)
		seatingSpots = newSeatingSpots
		newSeatingSpots = copySeatings(seatingSpots)
		println("one more iteration")
	}
	println("No change")
}

func printMe(seatingSpots [][]*spot) {
	height := getHeight(seatingSpots)
	width := getWidth(seatingSpots)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			spot := getSpot(seatingSpots, x, y)
			switch spot.state {
			case emptySeat:
				print("L")
			case occupied:
				print("#")
			case floor:
				print(".")
			}
		}
		println("")
	}
	println("---------")
}

func copySeatings(seatingSpots [][]*spot) [][]*spot {
	height := getHeight(seatingSpots)
	width := getWidth(seatingSpots)
	result := make([][]*spot, height)
	for y := 0; y < height; y++ {
		result[y] = make([]*spot, width)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			result[y][x] = dupSpot(getSpot(seatingSpots, x, y))
		}
	}
	return result
}

func iterateSeating(seatingSpots [][]*spot, newSeatingSpots *[][]*spot) bool {

	maxHeight := getHeight(seatingSpots)
	maxWidth := getWidth(seatingSpots)
	hasChanged := false

	for x := 0; x < maxWidth; x++ {
		for y := 0; y < maxHeight; y++ {
			//println(fmt.Sprintf("Checking %d,%d", x, y))
			if y == 1 && x == 0 {
				println("bingo")
			}

			current := getSpot(seatingSpots, x, y)
			newCurrentValue := getSpot(*newSeatingSpots, x, y)
			switch current.state {
			case emptySeat:
				if processEmptySeat(seatingSpots, current, x, y) {
					newCurrentValue.state = occupied
					hasChanged = true
				}
			case occupied:
				if processOccupiedSeat(seatingSpots, current, x, y) {
					newCurrentValue.state = emptySeat
					hasChanged = true
				}
			}
		}
	}
	return hasChanged
}

// Return true if state should change
func processEmptySeat(seatingSpots [][]*spot, current *spot, posX int, posY int) bool {
	atLeastOneOccupied := checkAtLeastOneAdjacent(seatingSpots, occupied, posX, posY)

	return !atLeastOneOccupied
}

func processOccupiedSeat(seatingSpots [][]*spot, current *spot, posX int, posY int) bool {
	isAtLeastFourOccupied := isAtLeastFourOccupied(seatingSpots, posX, posY)
	return isAtLeastFourOccupied
}

func checkAtLeastOneAdjacent(seatingSpots [][]*spot, stateValue int, posX int, posY int) bool {
	return countAdjacents(seatingSpots, stateValue, posX, posY) > 0
}

func isAtLeastFourOccupied(seatingSpots [][]*spot, posX int, posY int) bool {
	return countAdjacents(seatingSpots, occupied, posX, posY) >= 4
}

func countAdjacents(seatingSpots [][]*spot, stateValue int, posX int, posY int) int {

	adjacents := 0

	// Go W
	if move(seatingSpots, posX, posY, -1, 0, stateValue) {
		adjacents++
	}
	// Go E
	if move(seatingSpots, posX, posY, 1, 0, stateValue) {
		adjacents++
	}
	// Go N
	if move(seatingSpots, posX, posY, 0, -1, stateValue) {
		adjacents++
	}
	// Go S
	if move(seatingSpots, posX, posY, 0, 1, stateValue) {
		adjacents++
	}
	// Go NW
	if move(seatingSpots, posX, posY, -1, -1, stateValue) {
		adjacents++
	}
	// Go NE
	if move(seatingSpots, posX, posY, 1, -1, stateValue) {
		adjacents++
	}
	// Go SW
	if move(seatingSpots, posX, posY, -1, 1, stateValue) {
		adjacents++
	}
	// Go SE
	if move(seatingSpots, posX, posY, 1, 1, stateValue) {
		adjacents++
	}

	return adjacents
}

// Walk the seatingSpots until we found the expected state or the end of the seat chart
func move(seatingSpots [][]*spot, posX int, posY int, moveIncX int, moveIncY, expectedState int) bool {
	x := posX + moveIncX
	y := posY + moveIncY

	for isInside(seatingSpots, x, y) {
		seat := getSpot(seatingSpots, x, y)
		if isExpectedState(floor, seat) {
			x += moveIncX
			y += moveIncY
			continue
		}
		if isExpectedState(expectedState, seat) {
			return true
		}
		return false
	}
	return false
}

func isInside(seats [][]*spot, x int, y int) bool {
	return x >= 0 && y >= 0 && x < getWidth(seats) && y < getHeight(seats)
}

func isExpectedState(expected int, spot *spot) bool {
	return spot.state == expected
}

func getSpot(seatingSpots [][]*spot, posX int, posY int) *spot {
	return seatingSpots[posY][posX]
}

func getWidth(seatingSpots [][]*spot) int {
	return len(seatingSpots[0])
}

func getHeight(seatingSpots [][]*spot) int {
	return len(seatingSpots)
}
