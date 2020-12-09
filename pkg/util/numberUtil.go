package util

import "strconv"

func Convert(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}
