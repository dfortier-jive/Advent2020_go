package day13

import "testing"

func TestParseSchedule(t *testing.T) {
	busScheduleString := "7,13,x,x,59,x,31,19"
	busSchedule := parseBusSchedule(busScheduleString)

	if len(busSchedule) != 5 {
		t.Errorf("Found %d buses", len(busSchedule))
	}
	if busSchedule[0].id != 7 {
		t.Errorf("Unexpected bus id %d", busSchedule[0].id)
	}
	if busSchedule[1].id != 13 {
		t.Errorf("Unexpected bus id %d", busSchedule[1].id)
	}
	if busSchedule[2].id != 59 {
		t.Errorf("Unexpected bus id %d", busSchedule[2].id)
	}
	if busSchedule[3].id != 31 {
		t.Errorf("Unexpected bus id %d", busSchedule[3].id)
	}
	if busSchedule[4].id != 19 {
		t.Errorf("Unexpected bus id %d", busSchedule[4].id)
	}
}

func TestNextSchedule(t *testing.T) {
	busScheduleString := "7,59"
	busSchedule := parseBusSchedule(busScheduleString)
	busSeven := busSchedule[0]
	busFiftyNine := busSchedule[1]

	if busSeven.nextPickupWaitTime(0) != 0 {
		t.Errorf("Expected bus pickup %d but was %d", 0, busSeven.nextPickupWaitTime(0))
	}
	if busFiftyNine.nextPickupWaitTime(0) != 0 {
		t.Errorf("Expected bus pickup %d but was %d", 0, busFiftyNine.nextPickupWaitTime(0))
	}
	if busSeven.nextPickupWaitTime(7) != 0 {
		t.Errorf("Expected bus pickup %d but was %d", 0, busSeven.nextPickupWaitTime(7))
	}

	if busSeven.nextPickupWaitTime(1) != 6 {
		t.Errorf("Expected bus pickup %d but was %d", 6, busSeven.nextPickupWaitTime(1))
	}
	if busFiftyNine.nextPickupWaitTime(1) != 58 {
		t.Errorf("Expected bus pickup %d but was %d", 58, busFiftyNine.nextPickupWaitTime(1))
	}

	if busSeven.nextPickupWaitTime(9) != 5 {
		t.Errorf("Expected bus pickup %d but was %d", 5, busSeven.nextPickupWaitTime(9))
	}
	if busFiftyNine.nextPickupWaitTime(70) != 48 {
		t.Errorf("Expected bus pickup %d but was %d", 48, busFiftyNine.nextPickupWaitTime(70))
	}

}

func TestPart1Example(t *testing.T) {
	Part1(939, "7,13,x,x,59,x,31,19")
}

func TestPart1(t *testing.T) {
	Part1(1008169, "29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,653,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,17,x,x,x,x,x,23,x,x,x,x,x,x,x,823,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19")
}
