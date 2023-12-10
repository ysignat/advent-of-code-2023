package main

import (
	"advent/days"
	"flag"
	"fmt"
	"os"
)

type args struct {
	day  int
	part int
}

func main() {
	args := getArgs()

	var res string
	var err error

	switch args.day {
	case 1:
		switch args.part {
		case 1:
			res, err = days.Day1Part1()
		case 2:
			res, err = days.Day1Part2()
		default:
			fmt.Printf("Part %d for day %d doesn't exists\n", args.part, args.day)
			os.Exit(1)
		}
	case 2:
		switch args.part {
		case 1:
			res, err = days.Day2Part1()
		case 2:
			res, err = days.Day2Part2()
		default:
			fmt.Printf("Part %d for day %d doesn't exists\n", args.part, args.day)
			os.Exit(1)
		}
	default:
		fmt.Printf("Day %d doesn't exists\n", args.day)
		os.Exit(1)
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

func getArgs() args {
	var args args

	flag.IntVar(&args.day, "day", 1, "Specify the Advent of Code day")
	flag.IntVar(&args.part, "part", 1, "Specify the Advent of Code day part")
	flag.Parse()

	return args
}
