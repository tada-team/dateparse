package dateparse

import (
	"strings"
	"time"
)

var (
	seconds           = `сек|секунд|sec|seconds`
	minutes           = `мин|минут|минуту|min `
	hours             = `часов|hours|hour|часа|час`
	days              = `дней|дня|days`
	weeksWords        = `недель|неделю|недели|неделя|weeks|week`
	monthsWords       = `месяцев|месяца|month`
	years             = `лет|года|год|years`
	durationTimeWords = strings.Join([]string{seconds, minutes, hours, days, weeksWords, monthsWords, years}, "|")
)

var (
	durPrefix      = `(через|in|следующ[и]?[й]?[у]?[ю]?|следующий|следующую|next)`
	duration       = strings.Join([]string{today, tomorrow, afterTomorrow, afterAfterTomorrow, yesterday}, "|")
	durationTime   = `сек[у]?[н]?[д]?[а]?[у]?|мин[у]?[т]?[у]?[а]?[ы]?|min[u]?[t]?[e]?[s]?|час[о]?[в]?[а]?|hour[s]?|дн[е]?[й]?[я]?|day[s]?|недел[ь]?[я]?[и]?[ю]?|week[s]?|год|year[s]?|месяц[а]?[е]?[в]?|month[s]?`
	durationWds    = strings.Join([]string{duration, durationTime}, "|")
	durationSuffix = `(утр[а]?[о]?[м]?|morning|вечер[а]?[о]?[м]?|evening|\\|/|днем|полдень|полночь|midday|noon|midnight|ночью)`
)

var (
	one         = `one|один`
	two         = `two|два`
	three       = `three|три`
	four        = `four|четыре`
	five        = `five|пять`
	six         = `six|шесть`
	seven       = `seven|семь`
	eight       = `eight|восемь`
	nine        = `nine|девять`
	ten         = `ten|десять`
	wordNumbers = strings.Join([]string{one, two, three, four, five, six, seven, eight, nine, ten}, "|")
)

func checkWordNumber(s string) int64 {
	if v := forceInt64(s); v != 0 {
		return v
	}
	switch {
	case strings.Contains(one, s):
		return 1
	case strings.Contains(two, s):
		return 2
	case strings.Contains(three, s):
		return 3
	case strings.Contains(four, s):
		return 4
	case strings.Contains(five, s):
		return 5
	case strings.Contains(six, s):
		return 6
	case strings.Contains(seven, s):
		return 7
	case strings.Contains(eight, s):
		return 8
	case strings.Contains(nine, s):
		return 9
	case strings.Contains(ten, s):
		return 10
	}
	return 0
}

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
		v := checkWordNumber(bits[0])
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
