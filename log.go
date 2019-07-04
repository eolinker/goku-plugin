package goku_plugin

import (
	"strings"
	"time"
)

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

func ParsePeriod(period string)LogPeriod  {

	switch strings.ToLower( period ){
	case "hour":
		return PeriodHour
	case "day":
		return PeriodDay
	case "month":
		return PeriodMonth
	default:
		return PeriodHour
	}
}
func (p LogPeriod) String() string {
	switch p {
	case PeriodMonth:
		return "month"
	case PeriodDay:
		return "day"
	case PeriodHour:
		return "hour"
	default:
		return "unknown"
	}
}
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
