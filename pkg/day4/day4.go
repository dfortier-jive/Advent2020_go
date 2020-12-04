package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var validLabels = []string{"byr", "iyr", "eyr", "hgt", "pid", "ecl", "hcl"}
var northPoleId = []string{"cid"}

var validEcl = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

var rxHair = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var rxPassport = regexp.MustCompile(`^\d{9}$`)

type passport struct {
	fields map[string]string
}

func newPassport() passport {
	return passport{
		fields: make(map[string]string, 10),
	}
}

func readData() []passport {
	var f *os.File
	var err error
	if f, err = os.Open("passports.txt"); err != nil {
		panic("Unable to read file")
	}
	result := make([]passport, 0)
	passport := newPassport()

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			// Append to current passport
			ids := strings.Split(line, " ")
			for _, id := range ids {
				pair := strings.Split(id, ":")
				label := pair[0]
				value := pair[1]
				passport.fields[label] = value
			}
		} else {
			result = append(result, passport)
			// new passport
			passport = newPassport()
		}
	}
	result = append(result, passport)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

func (p *passport) ContainsRequiredFields() bool {
	for _, validLabel := range validLabels {
		if _, ok := p.fields[validLabel]; !ok {
			return false
		}
	}
	return true
}

func (p *passport) ContainsValidFields() bool {
	for _, validLabel := range validLabels {
		valid := true
		if fieldValue, ok := p.fields[validLabel]; ok {
			switch validLabel {
			case "byr":
				valid = validateInRange(convert(fieldValue), 1920, 2002)
			case "iyr":
				valid = validateInRange(convert(fieldValue), 2010, 2020)
			case "eyr":
				valid = validateInRange(convert(fieldValue), 2020, 2030)
			case "hgt":
				valid = validateHeight(fieldValue)
			case "hcl":
				valid = validateHairColor(fieldValue)
			case "ecl":
				valid = validateEyeColor(fieldValue)
			case "pid":
				valid = validatePassportNum(fieldValue)
			}
		} else {
			valid = false
		}
		if !valid {
			return false
		}
	}
	return true
}

func ValidateNbFields() {
	passports := readData()
	count := 0
	println(len(passports))
	for _, passport := range passports {
		if passport.ContainsRequiredFields() {
			count++
		}
	}
	println(fmt.Sprintf("Found %d valid passports", count))
}

func ValidateFieldValues() {
	passports := readData()
	count := 0
	for _, passport := range passports {
		if passport.ContainsValidFields() {
			count++
		}
	}
	println(fmt.Sprintf("Found %d valid passports", count))
}

func convert(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}

func validateInRange(value int, min int, max int) bool {
	return min <= value && value <= max
}

func validateHairColor(value string) bool {
	return rxHair.MatchString(value)
}

func validateEyeColor(value string) bool {
	for _, color := range validEcl {
		if value == color {
			return true
		}
	}
	return false
}

func validatePassportNum(value string) bool {
	return rxPassport.MatchString(value)
}

func validateHeight(value string) bool {
	if len(value) <= 2 {
		return false
	}
	unit := value[len(value)-2:]
	numValue := value[:len(value)-2]

	if unit == "in" {
		return validateInRange(convert(numValue), 59, 76)
	} else if unit == "cm" {
		return validateInRange(convert(numValue), 150, 193)
	} else {
		return false
	}
}
