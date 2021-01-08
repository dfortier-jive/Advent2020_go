package day13

import (
	"fmt"
	"math"
	"strings"

	"dfortier.org/advent2020/pkg/util"
)

type bus struct {
	id int
}

func newBus(id int) *bus {
	return &bus{
		id: id,
	}
}

func (b *bus) nextPickupWaitTime(timestamp int) int {
	remainder := timestamp % b.id
	if remainder == 0 {
		// Right on time!
		return 0
	}
	return b.id - remainder
}

func parseBusSchedule(busSchedule string) []*bus {
	allBuses := strings.Split(busSchedule, ",")
	buses := make([]*bus, 0)

	for _, busID := range allBuses {
		if busID != "x" {
			bus := newBus(util.Convert(busID))
			buses = append(buses, bus)
		}
	}
	return buses
}

func Part1(timeStamp int, busScheduleString string) {
	busSchedule := parseBusSchedule(busScheduleString)

	minWaitTime := math.MaxInt32
	var minWaitTimeBus *bus
	for _, bus := range busSchedule {
		nextPickupWait := bus.nextPickupWaitTime(timeStamp)
		if nextPickupWait < minWaitTime {
			minWaitTime = nextPickupWait
			minWaitTimeBus = bus
		}
	}
	println(fmt.Sprintf("Bus %d is next with %d wait time. Result %d",
		minWaitTimeBus.id, minWaitTime, minWaitTimeBus.id*minWaitTime))

}
