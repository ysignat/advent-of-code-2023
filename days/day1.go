package days

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Day1Part1() (string, error) {
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

func Day1Part2() (string, error) {
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
