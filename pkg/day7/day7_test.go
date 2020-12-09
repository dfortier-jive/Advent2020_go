package day7

import (
	"fmt"
	"testing"
)

func TestReadData(t *testing.T) {
	rules := readData()

	for k, v := range rules.bags {
		println(fmt.Sprintf("Node %s with %d child", k, len(v.subbags)))
	}

	t.Logf("Returned %d rules", len(rules.bags))
}

func TestVisitPart1(t *testing.T) {
	Day1VisitPart1()
}
