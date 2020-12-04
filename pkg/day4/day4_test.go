package day4

import (
	"testing"
)

func TestPassportValidationPart1(t *testing.T) {
	ValidateNbFields()
}

func TestPassportValidationPart2(t *testing.T) {
	ValidateFieldValues()
}

func TestMissingMultipleFields(t *testing.T) {
	passport := newPassport()
	passport.fields["ecl"] = "gry"
	passport.fields["pid"] = "gry"
	passport.fields["eyr"] = "gry"
	passport.fields["hcl"] = "gry"
	passport.fields["cid"] = "gry"
	//passport.fields["byr"] = "gry"
	//passport.fields["iyr"] = "gry"
	//passport.fields["hgt"] = "gry"
	if passport.ContainsRequiredFields() {
		t.Errorf("Passport should not be valid")
	}
}

func TestMissingCidOnly(t *testing.T) {
	passport := newPassport()
	passport.fields["ecl"] = "gry"
	passport.fields["pid"] = "gry"
	passport.fields["eyr"] = "gry"
	passport.fields["hcl"] = "gry"
	passport.fields["byr"] = "gry"
	passport.fields["iyr"] = "gry"
	passport.fields["hgt"] = "gry"
	//passport.fields["cid"] = "gry"

	if !passport.ContainsRequiredFields() {
		t.Errorf("Passport should be valid")
	}
}

func TestAllPresents(t *testing.T) {
	passport := newPassport()
	passport.fields["ecl"] = "gry"
	passport.fields["pid"] = "gry"
	passport.fields["eyr"] = "gry"
	passport.fields["hcl"] = "gry"
	passport.fields["byr"] = "gry"
	passport.fields["iyr"] = "gry"
	passport.fields["hgt"] = "gry"
	passport.fields["cid"] = "gry"

	if !passport.ContainsRequiredFields() {
		t.Errorf("Passport should be valid")
	}
}

func TestHeight(t *testing.T) {
	testCases := []struct {
		height      string
		expectValid bool
	}{
		{
			height:      "59in",
			expectValid: true,
		},
		{
			height:      "58in",
			expectValid: false,
		},
		{
			height:      "77in",
			expectValid: false,
		},
		{
			height:      "76in",
			expectValid: true,
		},
		{
			height:      "149cm",
			expectValid: false,
		},
		{
			height:      "150cm",
			expectValid: true,
		},
		{
			height:      "194cm",
			expectValid: false,
		},
		{
			height:      "193cm",
			expectValid: true,
		},
		{
			height:      "3pommes",
			expectValid: false,
		},
		{
			height:      "cm",
			expectValid: false,
		},
	}
	for _, c := range testCases {
		t.Log(c.height)
		if validateHeight(c.height) != c.expectValid {
			t.Errorf("Expected %s to be valid = %t", c.height, c.expectValid)
		}
	}
}

func TestHairColor(t *testing.T) {
	testCases := []struct {
		color       string
		expectValid bool
	}{
		{
			color:       "#123abc",
			expectValid: true,
		},
		{
			color:       "123abc",
			expectValid: false,
		},
		{
			color:       "#123abcx",
			expectValid: false,
		},
	}
	for _, c := range testCases {
		t.Log(c.color)
		if validateHairColor(c.color) != c.expectValid {
			t.Errorf("Expected %s to be valid = %t", c.color, c.expectValid)
		}
	}
}

func TestPassport(t *testing.T) {
	testCases := []struct {
		number      string
		expectValid bool
	}{
		{
			number:      "000000001",
			expectValid: true,
		},
		{
			number:      "0123456789",
			expectValid: false,
		},
	}
	for _, c := range testCases {
		t.Log(c.number)
		if validatePassportNum(c.number) != c.expectValid {
			t.Errorf("Expected %s to be valid = %t", c.number, c.expectValid)
		}
	}
}
