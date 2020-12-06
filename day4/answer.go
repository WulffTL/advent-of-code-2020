package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"regexp"
	"strconv"
	"strings"
)

type slope struct {
	rise int
	run  int
}

func main() {
	inputPath := "day4/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetValidPassportCount(input)
	fmt.Print(answer)
}

func GetValidPassportCount(input []string) int {
	validPassportCount := 0
	passports := getPassportsFromInput(input)
	for i := 0; i < len(passports); i++ {
		if isPassportValid(passports[i]) {
			validPassportCount++
		}
	}
	return validPassportCount
}

func getPassportsFromInput(input []string) [][]string {
	var passports [][]string
	var currentPassport []string
	currentPassportIndex := 0

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			passports = append(passports, currentPassport)
			currentPassport = nil
			currentPassportIndex++
			continue
		}

		fields := strings.Split(input[i], " ")
		currentPassport = append(currentPassport, fields...)

		if i == len(input)-1 {
			passports = append(passports, currentPassport)
		}

	}

	return passports
}

func isPassportValid(passport []string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	fieldsOnPassport := make(map[string]bool)
	for i := 0; i < len(passport); i++ {
		keyValue := strings.Split(passport[i], ":")
		fieldsOnPassport[keyValue[0]] = isFieldValid(keyValue[0], keyValue[1])
	}

	for i := 0; i < len(requiredFields); i++ {
		a := fieldsOnPassport[requiredFields[i]]
		if !a {
			return false
		}
	}

	return true
}

func isFieldValid(key string, value string) bool {
	switch key {
	case "byr":
		return isByrValid(value)
	case "iyr":
		return isIyrValid(value)
	case "eyr":
		return isEyrValid(value)
	case "hgt":
		return isHgtValid(value)
	case "hcl":
		return isHclValid(value)
	case "ecl":
		return isEclValid(value)
	case "pid":
		return isPidValid(value)
	default:
		return false
	}
}

func isByrValid(value string) bool {
	min := int64(1920)
	max := int64(2002)
	return isBetween(value, min, max)
}

func isIyrValid(value string) bool {
	min := int64(2010)
	max := int64(2020)
	return isBetween(value, min, max)
}

func isEyrValid(value string) bool {
	min := int64(2020)
	max := int64(2030)
	return isBetween(value, min, max)
}

func isHgtValid(value string) bool {
	if strings.Contains(value, "cm") {
		height := strings.Replace(value, "cm", "", 1)
		return isBetween(height, 150, 193)
	}

	if strings.Contains(value, "in") {
		height := strings.Replace(value, "in", "", 1)
		return isBetween(height, 59, 76)
	}

	return false
}

func isHclValid(value string) bool {
	if string(value[0]) != "#" {
		return false
	}
	for i := 1; i < len(value); i++ {
		matched, err := regexp.MatchString(`[0-9a-f]`, string(value[i]))
		if !matched || err != nil {
			return false
		}
	}

	return true
}

func isEclValid(value string) bool {
	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for i := 0; i < len(validEyeColors); i++ {
		if validEyeColors[i] == value {
			return true
		}
	}
	return false
}

func isPidValid(value string) bool {
	matched, err := regexp.MatchString(`^[0-9a-f]{9}$`, string(value))
	return matched && err == nil
}

func isBetween(value string, min int64, max int64) bool {
	v, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return false
	}

	return v >= min && v <= max
}
