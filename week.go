package dateparse

import (
	"strings"
	"time"
)

var (
	sunday    = `воскресенье|sunday`
	monday    = `понедельник|monday`
	tuesday   = `вторник|tuesday`
	wednesday = `среду|среда|wednesday`
	thursday  = `четверг|thursday`
	friday    = `пятницу|пятница|friday`
	saturday  = `субботу|суббота|saturday`
	weeks     = strings.Join([]string{
		sunday,
		monday,
		tuesday,
		wednesday,
		thursday,
		friday,
		saturday,
	}, "|")
)

var (
	shortSunday    = `вс|воскр|sun`
	shortMonday    = `пн|пнд|понед|mon`
	shortTuesday   = `вт|thu`
	shortWednesday = `ср|wed`
	shortThursday  = `чт|thu`
	shortFriday    = `пт|fri`
	shortSaturday  = `сб|sat`
	shortWeeks     = strings.Join([]string{
		shortSunday,
		shortMonday,
		shortTuesday,
		shortWednesday,
		shortThursday,
		shortFriday,
		shortSaturday,
	}, "|")
)

func parseWeekDay(s string, opts Opts) time.Time {
	date := getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), opts.TodayEndHour, 0, 0, opts)
	if weakDay := parseWeekDays(s); weakDay < 7 {
		v := weakDay - int(date.Weekday())
		if v < 0 {
			date = date.Add(time.Duration(v)*24*time.Hour + 7*24*time.Hour)
		} else {
			date = date.Add(time.Duration(v) * 24 * time.Hour)
		}
	}
	return date
}

func parseWeekDays(s string) int {
	switch {
	case strings.Contains(strings.Join([]string{sunday, shortSunday}, "|"), s):
		return 0
	case strings.Contains(strings.Join([]string{monday, shortMonday}, "|"), s):
		return 1
	case strings.Contains(strings.Join([]string{tuesday, shortTuesday}, "|"), s):
		return 2
	case strings.Contains(strings.Join([]string{wednesday, shortWednesday}, "|"), s):
		return 3
	case strings.Contains(strings.Join([]string{thursday, shortThursday}, "|"), s):
		return 4
	case strings.Contains(strings.Join([]string{friday, shortFriday}, "|"), s):
		return 5
	case strings.Contains(strings.Join([]string{saturday, shortSaturday}, "|"), s):
		return 6
	}
	return 7
}

func calculateWeekDuration(m []string, opts Opts, weekPosition int) (time.Time, string) {
	timePosition := weekPosition + 1
	if weekPosition < 0 {
		weekPosition = len(m) - 1
		timePosition = weekPosition - 2
	}
	date := parseWeekDay(m[weekPosition], opts)
	switch {
	case strings.Contains(durPrefix, m[weekPosition-1]):
		date = date.Add(24 * 7 * time.Hour)
	}
	if len(m) > 3 {
		switch {
		case strings.Contains(morning, m[timePosition]) && m[timePosition] != "":
			if date.Weekday() == opts.Now.Weekday() && opts.Now.Hour() > 10 {
				date = date.Add(24 * 7 * time.Hour)
			}
			return getDate(date.Year(), int(date.Month()), date.Day(), 10, 0, 0, opts), m[0]
		case strings.Contains(evening, m[timePosition]) && m[timePosition] != "":
			return date, m[0]
		case strings.Contains(noon, m[timePosition]) && m[timePosition] != "":
			return getDate(date.Year(), int(date.Month()), date.Day(), 12, 0, 0, opts), m[0]
		case strings.Contains(midnight, m[timePosition]) && m[timePosition] != "":
			return getDate(date.Year(), int(date.Month()), date.Day(), 0, 0, 0, opts), m[0]
		}
	}
	if date.Before(opts.Now) {
		date = date.Add(7 * 24 * time.Hour)
	}
	return date, m[0]
}
