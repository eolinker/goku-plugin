package goku_plugin

import "time"

type Logger interface {
	Log(string)
}
type LoggerGeneral interface {
	GenAccessLogger(dirName, fileName string, period LogPeriod) Logger
}

type LogPeriod int

const (
	PeriodMonth = iota
	PeriodDay
	PeriodHour
)

func (p LogPeriod) Format() string {
	t := time.Time{}
	t.Month()
	switch p {
	case PeriodMonth:
		return "2006-01-02-15"
	case PeriodDay:
		return "2006-01-02"
	case PeriodHour:
		return "2006-01-02-15"
	default:
		return "2006-01-02-15"
	}
}
