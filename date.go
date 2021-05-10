package dateparse

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	datePrefix = `(в|во|in|on|ровно|the|at)`
	dateSuffix = `(-ого|-го|-ва|-его|th|числа|date|\\|/|года|years|[.])`
	daySuffix  = `(-ого|-го|-ва|-его|th|числа|date|\\)`
)

var (
	baseDurRegex            = regexp.MustCompile(fmt.Sprintf(`(%s)[" "/]`, duration))
	baseDurOnlyRegex        = regexp.MustCompile(fmt.Sprintf(`(%s)$`, duration))
	baseDurTimeRegex        = regexp.MustCompile(fmt.Sprintf(`(\d\d?\d?)[" "](%s)`, durationTime))
	baseWeekOnlyRegex       = regexp.MustCompile(fmt.Sprintf(`^(%s|%s)$`, weeks, shortWeeks))
	baseWeekPrefixOnlyRegex = regexp.MustCompile(fmt.Sprintf(`^%s[" "](%s|%s)$`, datePrefix, weeks, shortWeeks))
	baseWeekPrefixRegex     = regexp.MustCompile(fmt.Sprintf(`^%s[" "](%s|%s)[" "]`, datePrefix, weeks, shortWeeks))
	baseWeekRegex           = regexp.MustCompile(fmt.Sprintf(`^(%s|%s)[" "]`, weeks, shortWeeks))
	weekDurSuffixRegex      = regexp.MustCompile(fmt.Sprintf(`%s[" "](%s)[" "]%s`, datePrefix, weeks, durationSuffix))
	durSuffixWeekRegex      = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[" "]%s[" "](%s)`, datePrefix, durationSuffix, datePrefix, weeks))
	durPrefixWeekRegex      = regexp.MustCompile(fmt.Sprintf(`%s[" "]%s[" "](%s)[" "]?%s?`, datePrefix, durPrefix, weeks, durationSuffix))
)

var (
	ddRegex          = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[" "]?%s`, datePrefix, dayDD, daySuffix))
	ddmmRegex        = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[/.]%s\s?%s?`, datePrefix, dayDD, monthMM, dateSuffix))
	ddMonthRegex     = regexp.MustCompile(fmt.Sprintf(`%s%s?[" "](%s)`, dayDD, dateSuffix, months))
	ddmmyyyyRegex    = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[/.]%s[/.]%s\s?%s?`, datePrefix, dayDD, monthMM, yearYYYY, dateSuffix))
	ddMonthyyyyRegex = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[" "/.](%s)[" "/.]%s\s?%s?`, datePrefix, dayDD, months, yearYYYY, dateSuffix))
	ddmmyyRegex      = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[/.]%s[/.]%s\s?%s?`, datePrefix, dayDD, monthMM, yearYY, dateSuffix))
	ddMonthyyRegex   = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[" "/.](%s)[" "/.]%s\s?%s`, datePrefix, dayDD, months, yearYY, dateSuffix))
	isoyyyymmddRegex = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[/.-]?%s[/.-]?%s\s?`, datePrefix, yearYYYY, monthMM, dayDD))
	isoyymmddRegex   = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[/.-]?%s[/.-]?%s`, datePrefix, yearYY, monthMM, dayDD))
)

var (
	durTimeRegex   = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[" "](\d\d?\d?)[" "]?(%s)?`, datePrefix, durPrefix, durationTime))
	durRegex       = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[" "](%s)?\s?(%s)`, datePrefix, durPrefix, wordNumbers, durationWds))
	wdsRegex       = regexp.MustCompile(fmt.Sprintf(`(%s)\b[" "/]?%s?`, durationWds, durationSuffix))
	wdsSuffuxRegex = regexp.MustCompile(fmt.Sprintf(`(%s)[" "/]%s[" "]%s`, durationWds, datePrefix, durationSuffix))
	wdsTimeRegex   = regexp.MustCompile(fmt.Sprintf(`%s[" "](\d\d)[" "](%s)`, datePrefix, hours))
)

var (
	mmddyyyyRegex = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[/.]%s[/.]%s\s?%s?`, datePrefix, monthMM, dayDD, yearYYYY, dateSuffix))
	mmddyyRegex   = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[/.]%s[/.]%s\s?%s?`, datePrefix, monthMM, dayDD, yearYY, dateSuffix))
	mmddRegex     = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[/.]%s\s?%s?`, datePrefix, monthMM, dayDD, dateSuffix))
)

func parseDate(s string, opts Opts) (t time.Time, st string) {
	switch {
	case baseDurOnlyRegex.MatchString(s):
		return calculateWordsDate(baseDurOnlyRegex.FindStringSubmatch(s), opts)
	case baseWeekPrefixOnlyRegex.MatchString(s):
		return calculateWeekDuration(baseWeekPrefixOnlyRegex.FindStringSubmatch(s), opts, 2)
	case baseWeekOnlyRegex.MatchString(s):
		return calculateWeekDuration(baseWeekOnlyRegex.FindStringSubmatch(s), opts, 1)
	case wdsSuffuxRegex.MatchString(s):
		return calculateWordsDate(wdsSuffuxRegex.FindStringSubmatch(s), opts)
	case baseDurRegex.MatchString(s):
		return calculateWordsDate(baseDurRegex.FindStringSubmatch(s), opts)
	case weekDurSuffixRegex.MatchString(s):
		return calculateWeekDuration(weekDurSuffixRegex.FindStringSubmatch(s), opts, 2)
	case baseWeekPrefixRegex.MatchString(s):
		return calculateWeekDuration(baseWeekPrefixRegex.FindStringSubmatch(s), opts, 2)
	case baseWeekRegex.MatchString(s):
		return calculateWeekDuration(baseWeekRegex.FindStringSubmatch(s), opts, 1)
	case durTimeRegex.MatchString(s):
		return calculateDuration(durTimeRegex.FindStringSubmatch(s), opts, 2)
	case durRegex.MatchString(s):
		return calculateDuration(durRegex.FindStringSubmatch(s), opts, 2)
	case durPrefixWeekRegex.MatchString(s):
		return calculateWeekDuration(durPrefixWeekRegex.FindStringSubmatch(s), opts, 3)
	case durSuffixWeekRegex.MatchString(s):
		return calculateWeekDuration(durSuffixWeekRegex.FindStringSubmatch(s), opts, -3)
	//case rareyyyymmdd.MatchString(s):
	//	return calculateFullDate(rareyyyymmdd.FindStringSubmatch(s), opts, 2, 3, 4)
	//case rareyymmdd.MatchString(s):
	//	return calculateFullDate(rareyymmdd.FindStringSubmatch(s), opts, 2, 3, 4)
	case ddMonthyyyyRegex.MatchString(s):
		return calculateFullDate(ddMonthyyyyRegex.FindStringSubmatch(s), opts, 4, 3, 2)
	case ddMonthyyRegex.MatchString(s):
		return calculateFullDate(ddMonthyyRegex.FindStringSubmatch(s), opts, 4, 3, 2)
	case ddmmyyyyRegex.MatchString(s):
		return calculateFullDate(ddmmyyyyRegex.FindStringSubmatch(s), opts, 4, 3, 2)
	case mmddyyyyRegex.MatchString(s):
		return calculateFullDate(mmddyyyyRegex.FindStringSubmatch(s), opts, 4, 2, 3)
	case ddmmyyRegex.MatchString(s):
		return calculateFullDate(ddmmyyRegex.FindStringSubmatch(s), opts, 4, 3, 2)
	case mmddyyRegex.MatchString(s):
		return calculateFullDate(mmddyyRegex.FindStringSubmatch(s), opts, 4, 2, 3)
	case isoyyyymmddRegex.MatchString(s):
		return calculateFullDate(isoyyyymmddRegex.FindStringSubmatch(s), opts, 2, 3, 4)
	case isoyymmddRegex.MatchString(s):
		return calculateFullDate(isoyymmddRegex.FindStringSubmatch(s), opts, 2, 3, 4)
	case ddMonthRegex.MatchString(s):
		return calculateDate(ddMonthRegex.FindStringSubmatch(s), opts, 3, 1)
	case ddmmRegex.MatchString(s):
		return calculateDate(ddmmRegex.FindStringSubmatch(s), opts, 3, 2)
	case mmddRegex.MatchString(s):
		return calculateDate(mmddRegex.FindStringSubmatch(s), opts, 2, 3)
	case wdsTimeRegex.MatchString(s):
		m := wdsTimeRegex.FindStringSubmatch(s)
		date := getDate(opts.Now.Year(), opts.Now.Month(), opts.Now.Day(), forceInt(m[2]), 0, 0, opts)
		if date.Before(opts.Now) {
			date = date.Add(24 * time.Hour)
		}
		return date, m[0]
	case baseDurTimeRegex.MatchString(s):
		return calculateDuration(baseDurTimeRegex.FindStringSubmatch(s), opts, 1)
	case wdsRegex.MatchString(s):
		return calculateWordsDate(wdsRegex.FindStringSubmatch(s), opts)
	case ddRegex.MatchString(s):
		m := ddRegex.FindStringSubmatch(s)
		day := forceInt(m[2])
		return getDate(opts.Now.Year(), opts.Now.Month(), day, opts.TodayEndHour, 0, 0, opts), m[0]
	}
	return opts.Now, st
}

func getDate(year int, month time.Month, day int, hour int, minute int, second int, opts Opts) time.Time {
	return time.Date(year, month, day, hour, minute, second, 0, opts.Now.Location())
}

func calculateDate(m []string, opts Opts, monthPosition int, dayPosition int) (time.Time, string) {
	month := opts.Now.Month()
	if mth := parseMonth(m[monthPosition]); mth != 0 {
		month = mth
	} else {
		month = time.Month(forceInt(m[monthPosition]))
	}

	year := opts.Now.Year()
	if month < opts.Now.Month() {
		year++
	}

	day := forceInt(m[dayPosition])
	return getDate(year, month, day, opts.TodayEndHour, 0, 0, opts), m[0]
}

func calculateFullDate(m []string, opts Opts, yearPosition int, monthPosition int, dayPosition int) (time.Time, string) {
	year := opts.Now.Year()
	if len(m[yearPosition]) == 2 {
		year = forceInt("20" + m[yearPosition][:2])
	} else {
		year = forceInt(m[yearPosition][:4])
	}
	date, _ := calculateDate(m, opts, monthPosition, dayPosition)
	if date.Month() < opts.Now.Month() && year == opts.Now.Year() {
		year += 1
	}
	return getDate(year, date.Month(), date.Day(), opts.TodayEndHour, 0, 0, opts), m[0]
}

func calculateWordsDate(m []string, opts Opts) (time.Time, string) {
	m = normalizeStrings(m)
	date := getDate(opts.Now.Year(), opts.Now.Month(), opts.Now.Day(), opts.Now.Hour(), opts.Now.Minute(), 0, opts)
	str := strings.Replace(m[0], "/", "", 1)
	switch {
	case strings.Contains(today, str) || strings.Contains(today, m[1]):
		date = getDate(opts.Now.Year(), opts.Now.Month(), opts.Now.Day(), opts.TodayEndHour, 0, 0, opts)
	case strings.Contains(tomorrow, str) || strings.Contains(tomorrow, m[1]):
		date = getDate(opts.Now.Year(), opts.Now.Month(), opts.Now.Day()+1, opts.TodayEndHour, 0, 0, opts)
	case strings.Contains(afterTomorrow, m[0]):
		date = getDate(opts.Now.Year(), opts.Now.Month(), opts.Now.Day()+2, opts.TodayEndHour, 0, 0, opts)
	case strings.Contains(afterAfterTomorrow, m[0]):
		date = getDate(opts.Now.Year(), opts.Now.Month(), opts.Now.Day()+3, opts.TodayEndHour, 0, 0, opts)
	case strings.Contains(yesterday, m[0]):
		date = getDate(opts.Now.Year(), opts.Now.Month(), opts.Now.Day()+365, opts.TodayEndHour, 0, 0, opts)
	}
	if len(m) > 2 {
		switch {
		case strings.Contains(morning, m[2]):
			date = getDate(date.Year(), date.Month(), date.Day(), 10, 0, 0, opts)
		case strings.Contains(noon, m[2]):
			date = getDate(date.Year(), date.Month(), date.Day(), 12, 0, 0, opts)
		case strings.Contains(evening, m[2]):
			date = getDate(date.Year(), date.Month(), date.Day(), 18, 0, 0, opts)
		case strings.Contains(midnight, m[2]):
			if date.Day() == opts.Now.Day() {
				date = date.Add(24 * time.Hour)
			}
			date = getDate(date.Year(), date.Month(), date.Day(), 0, 0, 0, opts)
		}
	}
	if len(m) > 3 {
		switch {
		case strings.Contains(noon, m[3]):
			date = getDate(date.Year(), date.Month(), date.Day(), 12, 0, 0, opts)
		case strings.Contains(midnight, m[3]):
			if date.Day() == opts.Now.Day() {
				date = date.Add(24 * time.Hour)
			}
			date = getDate(date.Year(), date.Month(), date.Day(), 0, 0, 0, opts)
		}
	}

	return date, m[0]
}
