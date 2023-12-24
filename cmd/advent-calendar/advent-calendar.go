package main

import (
	"flag"
	"fmt"
	"syscall"

	"github.com/ysignat/advent-of-code-2023/internal/calendar"
	"golang.org/x/term"
)

type args struct {
	year        uint
	day         uint
	part        uint
	datasetType string
	rootPath    string
}

func main() {
	args := getArgs()

	var dataset_storage calendar.DatasetGetter
	var err error
	if args.datasetType == "web" {
		fmt.Printf("You've chosen web dataset. Now, enter your session cookie for %s\n", calendar.ADVENT_OF_CODE_BASE_URL)
		sessionCookie, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}

		dataset_storage, err = calendar.NewWebDataset(string(sessionCookie))
		if err != nil {
			panic(err)
		}
	} else {
		dataset_storage, err = calendar.NewFilesystemDataset(args.rootPath)
		if err != nil {
			panic(err)
		}
	}

	calendar := calendar.Calendar{}

	day, err := calendar.GetDay(args.day)
	if err != nil {
		panic(err)
	}

	part, err := day.GetPart(args.part)
	if err != nil {
		panic(err)
	}

	dataset, err := dataset_storage.GetDataset(args.year, args.day)
	if err != nil {
		panic(err)
	}

	result, err := part.Solve(dataset)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result is: %s\n", result)
}

func getArgs() args {
	var args args

	flag.UintVar(&args.year, "year", 2023, "Advent of Code year")
	flag.UintVar(&args.day, "day", 1, "Advent of Code day")
	flag.UintVar(&args.part, "part", 1, "Advent of Code day part")
	flag.StringVar(
		&args.datasetType, "dataset-type", "filesystem",
		fmt.Sprintf(
			`Type of the dataset.
Possible values:
    - filesystem - datasets are stored on a local drive under the root path specified with -root-path. Files structure must be
        < -root-path >/
        └── < -year >/
            └── day< -day >
        needs -root-path to be set
    - web - datasets are taken directly from the Advent of Code website (%s)
        needs -session-cookie to be set
`,
			calendar.ADVENT_OF_CODE_BASE_URL,
		),
	)
	flag.StringVar(&args.rootPath, "root-path", "./",
		`Root path for a filesystem datasets storage.
Needed if -dataset-type filesystem specified`,
	)

	flag.Parse()

	return args
}
