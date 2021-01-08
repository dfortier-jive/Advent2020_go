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

func (b *bus) bigNextPickupWaitTime(timestamp int64) int {
	remainder := timestamp % int64(b.id)
	if remainder == 0 {
		// Right on time!
		return 0
	}
	return b.id - int(remainder)
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

var noBus = newBus(-1)

func Part2(busScheduleString string, startPoint int64) int64 {
	busSchedule := parseBusSchedulePart2(busScheduleString)

	var firstBusFrequency int64 = int64(busSchedule[0].id)
	var timestamp int64 = (startPoint / firstBusFrequency) * firstBusFrequency

	for ; ; timestamp += firstBusFrequency {
		if checkOtherBusesRespectRule(busSchedule, timestamp) {
			// found one!
			println(fmt.Sprintf("Timestamp is %d", timestamp))
			return timestamp
		}
	}
}

func checkOtherBusesRespectRule(busSchedule []*bus, currentTs int64) bool {
	// Position in the array is equal to the nb of minutes since the first bus
	for i, bus := range busSchedule[1:] {
		if bus != noBus {
			if !isExpectedNext(currentTs, i+1, bus) {
				// Not found, try another timestamp from firstBus
				return false
			}
		}
	}
	return true
}

func isExpectedNext(ts int64, expectedWait int, bus *bus) bool {
	pickupWait := bus.bigNextPickupWaitTime(ts)
	return pickupWait == expectedWait
}

func parseBusSchedulePart2(busSchedule string) []*bus {
	allBuses := strings.Split(busSchedule, ",")
	buses := make([]*bus, 0)

	for _, busID := range allBuses {
		var bus *bus
		if busID != "x" {
			bus = newBus(util.Convert(busID))
		} else {
			bus = noBus
		}
		buses = append(buses, bus)
	}
	return buses
}
