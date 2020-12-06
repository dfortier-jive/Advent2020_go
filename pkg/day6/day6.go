package day6

import (
	"bufio"
	"fmt"
	"os"
)

var questionID = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

type group struct {
	forms []formData
}

func newGroup() group {
	return group {
		forms: make([]formData, 0),
	}
}

func (g *group) getSameAnswers() int {
	count := 0
	nbForms := len(g.forms)
	for _, question := range questionID {
		countQuestion := 0
		for _, oneForm := range g.forms {
			if _, ok := oneForm.FieldsYes[question]; ok {
				countQuestion++
			}
		}
		if countQuestion == nbForms {
			// everyone answered yes for this one
			count++
		}
	}
	return count
}

func (g *group) getAllAnswers() int {
	count := 0
	for _, question := range questionID {
		countQuestion := 0
		for _, oneForm := range g.forms {
			if _, ok := oneForm.FieldsYes[question]; ok {
				countQuestion++
			}
		}
		if countQuestion > 0 {
			count++
		}
	}
	return count
}

type formData struct {
	FieldsYes map[rune]bool
}

func newFormData() formData {
	return formData{
		FieldsYes: make(map[rune]bool, 10),
	}
}

func (f *formData) getAnswers() []rune {
	result := make([]rune, 0)
	for k := range f.FieldsYes {
		result = append(result, k)
	}
	return result
}

func readData() []group {
	var f *os.File
	var err error
	if f, err = os.Open("customForm.txt"); err != nil {
		panic("Unable to read file")
	}
	result := make([]group, 0)

	oneGroup := newGroup()
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			// Append to current group
			onePerson := newFormData()
			for _, answer := range line {
				onePerson.FieldsYes[answer] = true
			}
			oneGroup.forms = append(oneGroup.forms, onePerson)
		} else {
			result = append(result, oneGroup)
			oneGroup = newGroup()
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
		count += oneGroup.getAllAnswers()
	}
	return count
}

func Part2() int {
	data := readData()
	count := 0
	for _, oneGroup := range data {
		count += oneGroup.getSameAnswers()
	}
	return count
}