package day6

import (
	"bufio"
	"fmt"
	"os"
)

type formData struct {
	FieldsYes map[rune]bool
}

func newFormData() formData {
	return formData{
		FieldsYes: make(map[rune]bool, 10),
	}
}

func (f *formData) getNbYes() int {
	return len(f.FieldsYes)
}

func (f *formData) getAnswers() []rune {
	result := make([]rune, 0)
	for k := range f.FieldsYes {
		result = append(result, k)
	}
	return result
}

func readData() []formData {
	var f *os.File
	var err error
	if f, err = os.Open("customForm.txt"); err != nil {
		panic("Unable to read file")
	}
	result := make([]formData, 0)

	oneGroup := newFormData()
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			// Append to current group
			for _, answer := range line {
				oneGroup.FieldsYes[answer] = true
			}
		} else {
			result = append(result, oneGroup)
			oneGroup = newFormData()
		}
	}
	result = append(result, oneGroup)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func Part1() int {
	data := readData()
	count := 0
	for _, oneGroup := range data {
		count += oneGroup.getNbYes()
	}
	return count
}