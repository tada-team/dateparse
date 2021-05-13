package dateparse

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var dateTimeRegex, _ = joinRegexp([]*regexp.Regexp{baseDurOnlyRegex, baseWeekOnlyRegex, baseWeekPrefixRegex, baseWeekPrefixOnlyRegex,
	baseDurRegex, baseWeekRegex, baseTimeOrientationRegex, durTimeRegex, baseDurTimeRegex, durRegex, wdsSuffuxRegex, wdsRegex,
	ddRegex, ddmmRegex, ddMonthRegex, ddmmyyyyRegex, mmddyyyyRegex, mmddRegex, ddMonthyyyyRegex, ddmmyyRegex, mmddyyRegex,
	ddMonthyyRegex, durPrefixWeekRegex, weekDurSuffixRegex, durSuffixWeekRegex, hhmmRegex, hhRegex, isoyyyymmddRegex, isoyymmddRegex, wdsTimeRegex}, "|")

func dateTimeParse(s string, opts Opts) (t time.Time, msg string) {
	if dateTimeRegex.MatchString(s) {

		marker := getMarker()
		s = strings.ReplaceAll(s, "://", marker)

		date, replacingDate := parseDate(s, opts)
		s = strings.Replace(s, strings.TrimSpace(replacingDate), "", 1)
		timeP, replacingTime := parseTime(s, opts)
		if (timeP.Before(opts.Now) || timeP == opts.Now) && date == opts.Now {
			date = date.Add(24 * time.Hour)
		}

		hour := timeP.Hour()
		minute := timeP.Minute()
		second := timeP.Second()

		replacingTime = strings.TrimSpace(replacingTime)
		if len(replacingTime) == 0 {
			hour = date.Hour()
			minute = date.Minute()
			second = date.Second()
		}

		s = strings.Replace(s, replacingTime, "", 1)
		s = strings.ReplaceAll(s, marker, "://")

		return getDate(date.Year(), date.Month(), date.Day(), hour, minute, second, opts), strings.TrimSpace(s)
	}
	return
}

func joinRegexp(regexps []*regexp.Regexp, sep string) (*regexp.Regexp, error) {
	var b strings.Builder
	for i, re := range regexps {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(re.String())
	}
	return regexp.Compile(b.String())
}

func getMarker() string {
	b := make([]byte, 20)
	rand.Read(b)
	return string(b)
}
