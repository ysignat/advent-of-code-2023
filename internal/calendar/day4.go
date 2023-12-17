package calendar

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type card struct {
	selected []int
	winning  []int
}

func newCard(s string) (card, error) {
	card_string_splitted := strings.Split(s, ":")
	numbers_string := card_string_splitted[1]
	numbers_string_splitted := strings.Split(numbers_string, "|")
	selected_numbers := numbers_string_splitted[0]
	winning_numbers := numbers_string_splitted[1]

	var card card
	for _, number_string := range strings.Split(selected_numbers, " ") {
		if number_string != "" {
			number, err := strconv.Atoi(number_string)
			if err != nil {
				return card, err
			}
			card.selected = append(card.selected, number)
		}
	}

	for _, number_string := range strings.Split(winning_numbers, " ") {
		if number_string != "" {
			number, err := strconv.Atoi(number_string)
			if err != nil {
				return card, err
			}
			card.winning = append(card.winning, number)
		}
	}

	return card, nil
}

func setIntersection(s1 map[int]bool, s2 map[int]bool) map[int]bool {
	s_intersection := map[int]bool{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for k := range s1 {
		if s2[k] {
			s_intersection[k] = true
		}
	}

	return s_intersection
}

func arrayIntoSet(arr []int) map[int]bool {
	set := map[int]bool{}
	for _, number := range arr {
		set[number] = true
	}

	return set
}

type day4Part1 struct{}

func (day4Part1) Solve() (string, error) {
	var sum float64

	data, err := getFileContent(4, 1)
	if err != nil {
		return fmt.Sprint(sum), err
	}

	splitted := strings.Split(string(data), "\n")

	for _, line := range splitted {
		card, err := newCard(line)
		if err != nil {
			return fmt.Sprint(sum), nil
		}

		selected_set := arrayIntoSet(card.selected)
		winning_set := arrayIntoSet(card.winning)

		intersection := setIntersection(selected_set, winning_set)
		intersections_count := len(intersection)

		if intersections_count != 0 {
			sum += math.Pow(2, float64(len(intersection)-1))
		}
	}

	return fmt.Sprint(sum), nil
}

type day4 struct{}

func (day4) GetPart(part int) (DayPartSolver, error) {
	switch part {
	case 1:
		return day4Part1{}, nil
	default:
		return nil, NoSuchPartError{part: part}
	}
}
