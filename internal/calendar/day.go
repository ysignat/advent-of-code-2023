package calendar

import "fmt"

type NoSuchPartError struct {
	day  int
	part int
}

func (e NoSuchPartError) Error() string {
	return fmt.Sprintf("part %d for day %d doesn't exist", e.part, e.day)
}

type DayPartGetter interface {
	GetPart(part int) (DayPartSolver, error)
}

type DayPartSolver interface {
	Solve() (string, error)
}
