package dateparse

import (
	"regexp"
	"strings"
	"time"
)

var dateTimeRegex, _ = joinRegexp([]*regexp.Regexp{baseDurOnlyRegex, baseWeekOnlyRegex, baseWeekPrefixRegex, baseWeekPrefixOnlyRegex,
	baseDurRegex, baseWeekRegex, baseTimeOrientationRegex, durTimeRegex, baseDurTimeRegex, durRegex, wdsSuffuxRegex, wdsRegex,
	ddRegex, ddmmRegex, ddMonthRegex, ddmmyyyyRegex, mmddyyyyRegex, mmddRegex, ddMonthyyyyRegex, ddmmyyRegex, mmddyyRegex,
	ddMonthyyRegex, durPrefixWeekRegex, weekDurSuffixRegex, durSuffixWeekRegex, hhmmRegex, hhRegex, isoyyyymmddRegex, isoyymmddRegex,
	rareyyyymmdd, rareyymmdd}, "|")

func dateTimeParse(s string, opts Opts) (t time.Time, msg string) {
	if dateTimeRegex.MatchString(s) {
		date, replacingDate := parseDate(s, opts)
		s = strings.Replace(s, strings.TrimSpace(replacingDate), "", 1)
		timeP, replacingTime := parseTime(s, opts)
		if (timeP.Before(opts.Now) || timeP == opts.Now) && date == opts.Now {
			date = date.Add(24 * time.Hour)
		}

		s = strings.Replace(s, strings.TrimSpace(replacingTime), "", 1)
		hour := timeP.Hour()
		minute := timeP.Minute()
		second := timeP.Second()
		if len(replacingTime) == 0 {
			hour = date.Hour()
			minute = date.Minute()
			second = date.Second()
		}
		return getDate(date.Year(), int(date.Month()), date.Day(), hour, minute, second, opts), strings.TrimSpace(s)
	}
	return
}

func joinRegexp(regexps []*regexp.Regexp, sep string) (*regexp.Regexp, error) {
	var result string
	var newSep string
	for _, re := range regexps {
		result += newSep + re.String()
		newSep = sep
	}
	return regexp.Compile(result)
}
