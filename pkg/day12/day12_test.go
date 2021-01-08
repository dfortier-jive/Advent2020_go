package day12

import (
	"fmt"
	"testing"
)

func TestChangeDirection(t *testing.T) {
	direction := newDirection(1, 0)
	if direction.incrementX != 1 || direction.incrementY != 0 {
		t.Errorf("New direction not facing east")
	}

	direction.changeDirectionUp(90)
	if direction.incrementX != 0 || direction.incrementY != 1 {
		t.Errorf("New direction not facing north")
	}

	direction.changeDirectionUp(90)
	if direction.incrementX != -1 || direction.incrementY != 0 {
		t.Errorf("New direction not facing west")
	}

	direction.changeDirectionUp(90)
	if direction.incrementX != 0 || direction.incrementY != -1 {
		t.Errorf("New direction not facing south")
	}

	direction.changeDirectionUp(90)
	if direction.incrementX != 1 || direction.incrementY != 0 {
		t.Errorf("New direction not facing east")
	}

	direction.changeDirectionUp(180)
	if direction.incrementX != -1 || direction.incrementY != 0 {
		t.Errorf("New direction not facing west")
	}

	direction.changeDirectionUp(270)
	if direction.incrementX != 0 || direction.incrementY != 1 {
		t.Errorf("New direction not facing north")
	}
	direction.changeDirectionDown(90)
	if direction.incrementX != 1 || direction.incrementY != 0 {
		t.Errorf("New direction not facing north")
	}
}

func TestChangeDirectionWaypoint(t *testing.T) {
	direction := newDirection(10, 1)
	direction.changeDirectionUp(90)
	if direction.incrementX != -1 || direction.incrementY != 10 {
		t.Errorf("New direction not facing %d, %d but %d, %d",
			-1, 10, direction.incrementX, direction.incrementY)
	}

	direction = newDirection(10, 1)
	direction.changeDirectionUp(180)
	if direction.incrementX != -10 || direction.incrementY != -1 {
		t.Errorf("New direction not facing %d, %d but %d, %d",
			-10, -1, direction.incrementX, direction.incrementY)
	}

	direction = newDirection(5, 10)
	direction.changeDirectionUp(270)
	if direction.incrementX != 10 || direction.incrementY != -5 {
		t.Errorf("New direction not facing %d, %d but %d, %d",
			10, -5, direction.incrementX, direction.incrementY)
	}
}

func TestPart1(t *testing.T) {
	ship := Part1()

	println(fmt.Sprintf("Final positon %d, %d. Manhattan distance is %d", ship.x, ship.y, ship.getManhattanDistance()))
}

func TestPart2(t *testing.T) {
	ship := Part2()

	println(fmt.Sprintf("Direction %d, %d", ship.direction.incrementX, ship.direction.incrementY))
	println(fmt.Sprintf("Final positon %d, %d. Manhattan distance is %d", ship.x, ship.y, ship.getManhattanDistance()))
}
