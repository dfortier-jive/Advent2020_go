package day3

import (
	"bufio"
	"fmt"
	"os"
)

func readData() [][]bool {
	var f *os.File
	var err error
	if f, err = os.Open("inputdata.txt"); err != nil {
		panic("Unable to read file")
	}
	result := make([][]bool, 323)

	scanner := bufio.NewScanner(bufio.NewReader(f))
	countLine := 0
	for scanner.Scan() {
		line := scanner.Text()
		result[countLine] = make([]bool, len(line))

		// Scan for trees
		for i, value := range line {
			if value == '#' {
				result[countLine][i] = true
			} else {
				result[countLine][i] = false
			}
		}

		countLine++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func IterateMaze(nbRight int, nbDown int) int {

	maze := readData()
	maxHeight := len(maze)
	maxWidth := len(maze[0])

	x := 0
	y := 0
	countTree := 0

	for y < maxHeight {
		println(fmt.Sprintf("Checking %d,%d", x, y))
		// is tree
		if maze[y][x] {
			countTree++
		}

		// move right
		x = (x + nbRight) % maxWidth
		y = y + nbDown
	}
	return countTree
}
