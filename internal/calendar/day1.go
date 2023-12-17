package calendar

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type day1Part1 struct{}

func (day1Part1) Solve() (string, error) {
	sum := 0

	data, err := getFileContent(1, 1)
	if err != nil {
		return fmt.Sprint(sum), err
	}

	splitted := strings.Split(string(data), "\n")
	for _, line := range splitted {
		var first_digit string
		var last_digit string

		for _, char := range line {
			if unicode.IsDigit(char) {
				digit := string(char)
				if first_digit == "" {
					first_digit = digit
				}
				last_digit = digit
			}

		}

		number, err := strconv.Atoi(first_digit + last_digit)
		if err != nil {
			return fmt.Sprint(sum), err
		}

		sum += number
	}

	return fmt.Sprint(sum), nil
}

type day1Part2 struct{}

func (day1Part2) Solve() (string, error) {
	string_digits := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var sum int

	data, err := getFileContent(1, 1)
	if err != nil {
		return fmt.Sprint(sum), err
	}

	splitted := strings.Split(string(data), "\n")
	for _, line := range splitted {
		var first_digit string
		var last_digit string

		for len(line) > 0 {
			for index, prefix := range string_digits {
				digit := fmt.Sprint(index + 1)
				if strings.HasPrefix(line, prefix) || strings.HasPrefix(line, digit) {
					if first_digit == "" {
						first_digit = digit
					}
					last_digit = digit

					break
				}
			}

			line = line[1:]
		}

		number, err := strconv.Atoi(first_digit + last_digit)
		if err != nil {
			return fmt.Sprint(sum), err
		}

		sum += number
	}

	return fmt.Sprint(sum), nil
}

type day1 struct{}

func (day1) GetPart(part int) (DayPartSolver, error) {
	switch part {
	case 1:
		return day1Part1{}, nil
	case 2:
		return day1Part2{}, nil
	default:
		return nil, NoSuchPartError{part: part}
	}
}
