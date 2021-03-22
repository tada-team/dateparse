package dateparse

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	timePrefix = `(с|в|к|by|at)`
	timeSuffix = `(утра|вечера|мин[у]?[т]?|a[.]?m|p[.]?m)`
)

var (
	hhmmRegex                = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[.:]%s`, timePrefix, hourHH, minuteMM))
	hhRegex                  = regexp.MustCompile(fmt.Sprintf(`%s[" "]%s\s?%s?`, timePrefix, hourHH, timeSuffix))
	baseTimeOrientationRegex = regexp.MustCompile(fmt.Sprintf(`%s?%s?[" "]?%s`, timePrefix, datePrefix, durationSuffix))
)

func parseTime(s string, opts Opts) (t time.Time, st string) {
	switch {
	case hhmmRegex.MatchString(s):
		return calculateTime(hhmmRegex.FindStringSubmatch(s), opts)
	case hhRegex.MatchString(s):
		return calculateTime(hhRegex.FindStringSubmatch(s), opts)
	case baseTimeOrientationRegex.MatchString(s):
		return calculateTime(baseTimeOrientationRegex.FindStringSubmatch(s), opts)
	}
	return opts.Now, st
}

func calculateTime(t []string, opts Opts) (time.Time, string) {
	m := forceList(strings.Join(t[1:], ","))
	switch len(m) {
	case 4:
		hour := forceInt(m[2])
		minute := forceInt(m[3])
		return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), hour, minute, 0, opts), t[0]
	case 3:
		hour := forceInt(m[1])
		switch {
		case strings.Contains(morning, m[2]):
			if hour > 12 {
				hour -= 12
			}
			return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), hour, 0, 0, opts), t[0]
		case strings.Contains(evening, m[2]):
			if hour < 12 {
				hour += 12
			}
			return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), hour, 0, 0, opts), t[0]
		}
		minute := forceInt(m[2])
		return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), hour, minute, 0, opts), t[0]
	case 2:
		switch {
		case strings.Contains(morning, m[1]):
			return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), 10, 0, 0, opts), t[0]
		case strings.Contains(noon, m[1]):
			return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), 12, 0, 0, opts), t[0]
		case strings.Contains(timePrefix, m[0]):
			return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), forceInt(m[1]), 0, 0, opts), t[0]
		}
		return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), forceInt(m[0]), forceInt(m[1]), 0, opts), t[0]
	case 1:
		switch {
		case strings.Contains(morning, m[0]):
			return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), 10, 0, 0, opts), t[0]
		case strings.Contains(evening, m[0]):
			return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), 18, 0, 0, opts), t[0]
		case strings.Contains(noon, m[0]):
			return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), 12, 0, 0, opts), t[0]
		case strings.Contains(midnight, m[0]):
			return getDate(opts.Now.Year(), int(opts.Now.Month()), opts.Now.Day(), 0, 0, 0, opts), t[0]

		}
	}
	return opts.Now, t[0]
}
