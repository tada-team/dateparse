package dateparse

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	timePrefix = `(в|к|by|at)`
	timeSuffix = `(утра|вечера|мин[у]?[т]?|a[.]?m|p[.]?m)`
)

var (
	HHMMRegex                = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s[.:]%s`, timePrefix, HH, MM))
	HHRegex                  = regexp.MustCompile(fmt.Sprintf(`%s[" "]%s\s?%s?`, timePrefix, HH, timeSuffix))
	BaseTimeOrientationRegex = regexp.MustCompile(fmt.Sprintf(`%s?[" "]?%s`, datePrefix, durationSuffix))
)

func parseTime(s string, opts Opts) (t time.Time, st string) {
	switch {
	case HHMMRegex.MatchString(s):
		return calculateTime(HHMMRegex.FindStringSubmatch(s), opts)
	case HHRegex.MatchString(s):
		return calculateTime(HHRegex.FindStringSubmatch(s), opts)
	case BaseTimeOrientationRegex.MatchString(s):
		return calculateTime(BaseTimeOrientationRegex.FindStringSubmatch(s), opts)
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
