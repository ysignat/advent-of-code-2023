package calendar

import "fmt"

type NoSuchPartError struct {
	day  uint
	part uint
}

func (e NoSuchPartError) Error() string {
	return fmt.Sprintf("part %d for day %d doesn't exist", e.part, e.day)
}

type DayPartGetter interface {
	GetPart(part uint) (DayPartSolver, error)
}

type DayPartSolver interface {
	Solve(dataset string) (string, error)
}
