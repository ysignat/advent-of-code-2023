package calendar

import "fmt"

type DayGetter interface {
	GetDay(day int) (DayPartGetter, error)
}

type NoSuchDayError struct {
	day int
}

func (e NoSuchDayError) Error() string {
	return fmt.Sprintf("day %d doesn't exist", e.day)
}

type Calendar struct{}

func (Calendar) GetDay(day int) (DayPartGetter, error) {
	switch day {
	case 1:
		return day1{}, nil
	case 2:
		return day2{}, nil
	case 4:
		return day4{}, nil
	default:
		return nil, NoSuchDayError{day: day}
	}

}
