package day12

import (
	"fmt"
	"testing"
)

func TestChangeDirection(t *testing.T) {
	direction := newDirection()
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
}

func TestPart1(t *testing.T) {
	ship := Part1()

	println(fmt.Sprintf("Final positon %d, %d. Manhattan distance is %d", ship.x, ship.y, ship.getManhattanDistance()))
}
