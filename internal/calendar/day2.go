package calendar

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	RED_COUNT   = 12
	GREEN_COUNT = 13
	BLUE_COUNT  = 14
)

type set struct {
	red_count   int
	blue_count  int
	green_count int
}

type game struct {
	id   int
	sets []set
}

func newSet(s string) (set, error) {
	set_map := make(map[string]int)
	splitted := strings.Split(s, ", ")

	for _, color_value := range splitted {
		color_value_splitted := strings.Split(color_value, " ")

		value, err := strconv.Atoi(color_value_splitted[0])
		if err != nil {
			return set{}, err
		}

		color := color_value_splitted[1]

		set_map[color] = value
	}

	return set{red_count: set_map["red"], blue_count: set_map["blue"], green_count: set_map["green"]}, nil
}

func newGame(s string) (game, error) {
	game_string_splitted := strings.Split(s, ": ")
	game_id_string := game_string_splitted[0]
	id_string, _ := strings.CutPrefix(game_id_string, "Game ")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		return game{}, err
	}

	var sets []set
	sets_string := game_string_splitted[1]
	sets_string_splitted := strings.Split(sets_string, "; ")

	for _, set_string := range sets_string_splitted {
		set, err := newSet(set_string)
		if err != nil {
			return game{}, nil
		}
		sets = append(sets, set)
	}

	return game{id: id, sets: sets}, nil

}

type day2Part1 struct{}

func (day2Part1) Solve(dataset string) (string, error) {
	sum := 0

	splitted := strings.Split(dataset, "\n")

games_loop:
	for _, line := range splitted {
		game, err := newGame(line)
		if err != nil {
			return fmt.Sprint(sum), nil
		}

		for _, set := range game.sets {
			if set.blue_count > BLUE_COUNT || set.green_count > GREEN_COUNT || set.red_count > RED_COUNT {
				continue games_loop
			}
		}

		sum += game.id
	}

	return fmt.Sprint(sum), nil
}

type day2Part2 struct{}

func (day2Part2) Solve(dataset string) (string, error) {
	sum := 0

	splitted := strings.Split(string(dataset), "\n")

	for _, line := range splitted {
		game, err := newGame(line)
		if err != nil {
			return fmt.Sprint(sum), nil
		}

		var min_red int
		var min_blue int
		var min_green int

		for _, set := range game.sets {
			if set.red_count > min_red {
				min_red = set.red_count
			}
			if set.green_count > min_green {
				min_green = set.green_count
			}
			if set.blue_count > min_blue {
				min_blue = set.blue_count
			}
		}

		sum += min_blue * min_green * min_red
	}

	return fmt.Sprint(sum), nil
}

type day2 struct{}

func (day2) GetPart(part uint) (DayPartSolver, error) {
	switch part {
	case 1:
		return day2Part1{}, nil
	case 2:
		return day2Part2{}, nil
	default:
		return nil, NoSuchPartError{part: part}
	}
}
