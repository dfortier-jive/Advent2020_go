package day5

import(
	"bufio"
	"fmt"
	"os"
	"sort"
)

const nbRows = 128
const nbColums = 8

func readData() []string {
	var f *os.File
	var err error
	if f, err = os.Open("boardingPass.txt"); err != nil {
		panic("Unable to read file")
	}
	result := make([]string, 0)

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func processPass(pass string) (int, int) {
	maxRow := nbRows - 1
	minRow := 0
	maxCol := nbColums - 1
	minCol := 0
	
	for i, char := range pass {
		if (i < 7) {
			switch char {
			case 'F':
				maxRow = minRow + ((maxRow - minRow) / 2)
			case 'B':
				minRow = maxRow - ((maxRow - minRow)/ 2)
			}
		} else {
			switch char {
			case 'R':
				minCol = maxCol - ((maxCol - minCol)/ 2)
			case 'L':
				maxCol = minCol + ((maxCol - minCol) / 2)
			}
		}
	}
	return maxRow, maxCol
}

func getSeatId(row int, col int) int {
	return row * 8 + col
}

func Part1() {
	// readData
	data := readData()
	max := 0

	for _, pass := range data {
		row, col := processPass(pass)
		id := getSeatId(row, col)
		println(fmt.Sprintf("Passport %s seat is %d, %d (id: %d)", pass, row, col, id))
		if id > max {
			max = id
		}
	}
	println(fmt.Sprintf("Highest seat is %d", max))
}

func findMissingId(data []string) int {
	allIds := make([]int, 0)

	for _, pass := range data {
		row, col := processPass(pass)
		id := getSeatId(row, col)
		println(fmt.Sprintf("Passport %s seat is %d, %d (id: %d)", pass, row, col, id))
		allIds = append(allIds, id)
	}
	sort.Ints(allIds)
	prec := allIds[0] - 1

	for i := 0; i < len(allIds); i++ {
		current := allIds[i]
		if current != prec + 1 {
			return current - 1
		}
		prec = current
	}
	return -1
}

func Part2() {
	data := readData()
	missing := findMissingId(data)

	println(fmt.Sprintf("Missing seat is %d", missing))

}