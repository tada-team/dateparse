package dateparse

import (
	"strings"
	"time"
)

var (
	seconds           = `сек|секунд|sec|seconds`
	minutes           = `мин|минут|минуту|min `
	hours             = `час|часов|hour|часа`
	days              = `дней|дня|days`
	weeksWords        = `недель|неделю|недели|неделя|weeks|week`
	monthsWords       = `месяцев|месяца|month`
	years             = `лет|года|год|years`
	durationTimeWords = strings.Join([]string{seconds, minutes, hours, days, weeksWords, monthsWords, years}, "|")
)

var (
	durPrefix      = `(через|in|следующ[и]?[й]?[у]?[ю]?|следующий|next)`
	duration       = strings.Join([]string{today, tomorrow, afterTomorrow, afterAfterTomorrow, yesterday}, "|")
	durationTime   = `сек[у]?[н]?[д]?[а]?[у]?|мин[у]?[т]?[у]?[а]?[ы]?|min[u]?[t]?[e]?[s]?|час[о]?[в]?[а]?|hour[s]?|дн[е]?[й]?[я]?|day[s]?|недел[ь]?[я]?[и]?[ю]?|week[s]?|год|year[s]?|месяц[а]?[е]?[в]?|month[s]?`
	durationWds    = strings.Join([]string{duration, durationTime}, "|")
	durationSuffix = `(утр[а]?[о]?[м]?|morning|вечер[а]?[о]?[м]?|evening|\\|/|днем|полдень|полночь|midday|noon|midnight|ночью)`
)

func calculateDuration(m []string, opts Opts, k int) (time.Time, string) {
	dur := durationParse(m[k:], opts)
	if dur > 0 {
		return opts.Now.Add(dur).In(opts.Now.Location()), m[0]
	}
	return opts.Now, m[0]
}

func durationParse(bits []string, opts Opts) (dur time.Duration) {
	if strings.Contains(durPrefix, bits[0]) {
		return durationParse(forceList(strings.Join(bits[1:], ",")), opts)
	}

	switch len(bits) {
	case 1:
		word := bits[0]
		switch {
		case strings.Contains(durationTimeWords, word):
			return durationParse([]string{"1", word}, opts)
		}
		if forceInt64(word) > 0 {
			return durationParse([]string{word, "минут"}, opts)
		}
	case 2:
		v := forceInt64(bits[0])
		if v == 0 {
			return
		}
		word := strings.TrimSpace(bits[1])
		switch {
		case strings.Contains(seconds, word):
			return time.Duration(v) * time.Second
		case strings.Contains(minutes, word):
			return time.Duration(v) * time.Minute
		case strings.Contains(hours, word):
			return time.Duration(v) * time.Hour
		case strings.Contains(days, word):
			return time.Duration(v) * time.Hour * 24
		case strings.Contains(weeksWords, word):
			return time.Duration(v) * time.Hour * 24 * 7
		case strings.Contains(monthsWords, word):
			return time.Duration(v) * time.Hour * 24 * 31 // XXX:
		case strings.Contains(years, word):
			return time.Duration(v) * time.Hour * 24 * 365 // XXX:
		default:
			return durationParse(bits[:1], opts)

		}
	}
	return
}
