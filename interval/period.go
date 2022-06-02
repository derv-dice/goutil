package interval

import (
	"time"
)

type PeriodKind int

const (
	Nanosecond PeriodKind = iota + 1
	Microsecond
	Millisecond
	Second
	Minute
	Hour
	Day
	Week
	Month
	Year
)

type Period struct {
	kind      PeriodKind
	dur       time.Duration
	startFrom time.Time // Used only for Month and Year
}

func (p *Period) Duration() time.Duration {
	return p.dur
}

func NewPeriodNano(factor int) *Period {
	return newPeriod(Nanosecond, factor, true, time.Time{})
}

func NewPeriodMicro(factor int) *Period {
	return newPeriod(Microsecond, factor, true, time.Time{})
}

func NewPeriodMilli(factor int) *Period {
	return newPeriod(Millisecond, factor, true, time.Time{})
}

func NewPeriodSec(factor int) *Period {
	return newPeriod(Millisecond, factor, true, time.Time{})
}

func NewPeriodMin(factor int) *Period {
	return newPeriod(Minute, factor, true, time.Time{})
}

func NewPeriodHour(factor int) *Period {
	return newPeriod(Hour, factor, true, time.Time{})
}

func NewPeriodDay(factor int) *Period {
	return newPeriod(Hour, factor, true, time.Time{})
}

func NewPeriodWeek(factor int) *Period {
	return newPeriod(Week, factor, true, time.Time{})
}

func NewPeriodMonth(factor int, startDate time.Time) *Period {
	return newPeriod(Month, factor, false, startDate)
}

func NewPeriodYear(factor int, startDate time.Time) *Period {
	return newPeriod(Year, factor, false, startDate)
}

func newPeriod(kind PeriodKind, factor int, isDiscrete bool, startDate time.Time) (p *Period) {
	startDate = startDate.UTC()

	if factor <= 0 {
		factor = 1
	}

	p = &Period{
		kind: kind,
	}

	switch kind {
	case Nanosecond:
		p.dur = time.Nanosecond
	case Microsecond:
		p.dur = time.Microsecond
	case Millisecond:
		p.dur = time.Millisecond
	case Second:
		p.dur = time.Second
	case Minute:
		p.dur = time.Minute
	case Hour:
		p.dur = time.Hour
	case Day:
		p.dur = time.Hour * 24
	case Week:
		p.dur = time.Hour * 168
	case Month:
		tmp := startDate.AddDate(0, factor, 0)
		p.dur = tmp.Sub(startDate)
	case Year:
		tmp := startDate.AddDate(factor, 0, 0)
		p.dur = tmp.Sub(startDate)
	}

	if isDiscrete {
		p.dur *= time.Duration(factor)
	}

	return
}
