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
func newDirection(incX int, incY int) *direction {
	return &direction{
		incrementX: incX,
		incrementY: incY,
	}
}

type ship struct {
	x         int
	y         int
	direction *direction
}

func newShip() *ship {
	return newShipWithDirection(newDirection(1, 0))
}

func newShipWithDirection(direction *direction) *ship {
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

func (s *ship) MovePart2(instruction rune, value int) {
	println(fmt.Sprintf("Executing %#U to %d. Ship position before %d, %d. Waypoint vector %d, %d",
		instruction, value, s.x, s.y, s.direction.incrementX, s.direction.incrementY))
	switch instruction {
	case 'F':
		s.goForward(value)
	case 'N':
		s.direction.changeDirectionCoord(0, value)
	case 'S':
		s.direction.changeDirectionCoord(0, -1*value)
	case 'E':
		s.direction.changeDirectionCoord(value, 0)
	case 'W':
		s.direction.changeDirectionCoord(-1*value, 0)
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

func (d *direction) changeDirectionUp(angle int) {
	if angle%90 != 0 {
		panic(fmt.Sprintf("Angle is not a multiple of 90 degres. Found %d", angle))
	}
	floatAngle := float64(angle)
	floatRadians := floatAngle * math.Pi / 180

	x2 := float64(d.incrementX)*math.Cos(floatRadians) - float64(d.incrementY)*math.Sin(floatRadians)
	y2 := float64(d.incrementX)*math.Sin(floatRadians) + float64(d.incrementY)*math.Cos(floatRadians)

	d.incrementX = int(math.Round(x2))
	d.incrementY = int(math.Round(y2))
}

func (d *direction) changeDirectionDown(angle int) {
	d.changeDirectionUp(360 - angle)
}

func (d *direction) changeDirectionCoord(incX int, incY int) {
	d.incrementX += incX
	d.incrementY += incY
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

func Part2() *ship {
	instructions := readInstructions()
	ship := newShipWithDirection(newDirection(10, 1))

	for _, instruction := range instructions {
		ship.MovePart2(instruction.instruction, instruction.value)
	}

	return ship
}
