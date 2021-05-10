package dateparse

import (
	"strings"
	"time"
)

type Opts struct {
	TodayEndHour int
	Now          time.Time
}

func Parse(s string, opts *Opts) (time.Time, string) {
	if opts == nil {
		opts = new(Opts)
	}
	if opts.TodayEndHour == 0 {
		opts.TodayEndHour = 18
	}
	date, msg := dateTimeParse(strings.TrimSpace(strings.ToLower(s)), *opts)
	return date.Round(time.Second), msg
}
