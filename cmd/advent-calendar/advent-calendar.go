package main

import (
	"flag"
	"fmt"

	"github.com/ysignat/advent-of-code-2023/internal/calendar"
)

type args struct {
	day  int
	part int
}

func main() {
	args := getArgs()

	calendar := calendar.Calendar{}

	day, err := calendar.GetDay(args.day)
	if err != nil {
		panic(err)
	}

	part, err := day.GetPart(args.part)
	if err != nil {
		panic(err)
	}

	result, err := part.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result is: %s\n", result)
}

func getArgs() args {
	var args args

	flag.IntVar(&args.day, "day", 1, "Specify the Advent of Code day")
	flag.IntVar(&args.part, "part", 1, "Specify the Advent of Code day part")
	flag.Parse()

	return args
}
