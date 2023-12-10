package main

import (
	"advent/days"
	"flag"
	"fmt"
	"os"
)

type Args struct {
	day  int
	part int
}

func main() {
	args := getArgs()

	var res string
	var err error

	if args.day == 1 {
		if args.part == 1 {
			res, err = days.Day1Part1()
		} else if args.part == 2 {
			res, err = days.Day1Part2()
		}
	}

	if err == nil {
		fmt.Printf("Solution for day %d part %d problem is:\n", args.day, args.part)
		fmt.Println(res)
		os.Exit(0)
	} else {
		fmt.Printf("Unexpected error while solving day %d part %d problem\n", args.day, args.part)
		panic(err)
	}

}

func getArgs() Args {
	var args Args

	flag.IntVar(&args.day, "day", 1, "Specify the Advent of Code day")
	flag.IntVar(&args.part, "part", 1, "Specify the Advent of Code day part")
	flag.Parse()

	return args
}
