package day12

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"dfortier.org/advent2020/pkg/util"
)

type direction struct {
	incrementX int
	incrementY int
}

// Face East
func newDirection() *direction {
	return &direction{
		incrementX: 1,
		incrementY: 0,
	}
}

type ship struct {
	x         int
	y         int
	direction *direction
}

func newShip() *ship {
	direction := newDirection()
	return &ship{
		x:         0,
		y:         0,
		direction: direction,
	}
}

type instruction struct {
	instruction rune
	value       int
}

func newInstruction(instructionRune rune, value int) instruction {
	return instruction{
		instruction: instructionRune,
		value:       value,
	}
}

func (s *ship) Move(instruction rune, value int) {
	println(fmt.Sprintf("Executing %#U to %d. Ship position before %d, %d", instruction, value, s.x, s.y))
	switch instruction {
	case 'F':
		s.goForward(value)
	case 'N':
		s.changePosition(0, value)
	case 'S':
		s.changePosition(0, -1*value)
	case 'E':
		s.changePosition(value, 0)
	case 'W':
		s.changePosition(-1*value, 0)
	case 'L':
		s.direction.changeDirectionUp(value)
	case 'R':
		s.direction.changeDirectionDown(value)
	}
}

func (s *ship) goForward(value int) {
	s.changePosition(value*s.direction.incrementX, value*s.direction.incrementY)
}

func (s *ship) changePosition(incrementX int, incrementY int) {
	s.x = s.x + incrementX
	s.y = s.y + incrementY
}

func (s *ship) getManhattanDistance() int {
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func (s *direction) changeDirectionUp(angle int) {
	if angle%90 != 0 {
		panic(fmt.Sprintf("Angle is not a multiple of 90 degres. Found %d", angle))
	}
	floatAngle := float64(angle)
	floatRadians := floatAngle * math.Pi / 180

	x2 := float64(s.incrementX)*math.Cos(floatRadians) - float64(s.incrementY)*math.Sin(floatRadians)
	y2 := float64(s.incrementX)*math.Sin(floatRadians) + float64(s.incrementY)*math.Cos(floatRadians)

	s.incrementX = int(x2)
	s.incrementY = int(y2)
}

func (s *direction) changeDirectionDown(angle int) {
	s.changeDirectionUp(360 - angle)
}

func readInstructions() []instruction {
	var f *os.File
	var err error
	var result = make([]instruction, 0)
	if f, err = os.Open("instructions.txt"); err != nil {
		panic("Unable to read file")
	}

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()

		instructionRune := rune(line[0])
		value := util.Convert(line[1:])
		instruction := newInstruction(instructionRune, value)

		result = append(result, instruction)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func Part1() *ship {
	instructions := readInstructions()
	ship := newShip()

	for _, instruction := range instructions {
		ship.Move(instruction.instruction, instruction.value)
	}

	return ship
}
