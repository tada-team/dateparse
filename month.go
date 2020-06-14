package dateparse

import "strings"

var (
	january   = `января|янв|january|jan`
	february  = `февраля|фев|february|feb`
	march     = `марта|мар|march|mar`
	april     = `апреля|апрель|апр|april|apr`
	may       = `мая|май|may`
	june      = `июня|июнь|june|jun`
	july      = `июля|июль|july|jul`
	august    = `августа|август|august|aug`
	september = `сентября|сентябрь|сент|september|sep`
	october   = `октября|октябрь|october|oct`
	november  = `ноября|ноябрь|november|nov`
	december  = `декабря|дек|december|dec`
	months    = strings.Join([]string{january, february, march, april, may, june, july, august, september, october, november, december}, "|")
)

func parseMonth(monthStr string) int {
	switch {
	case strings.Contains(january, monthStr):
		return 1
	case strings.Contains(february, monthStr):
		return 2
	case strings.Contains(march, monthStr):
		return 3
	case strings.Contains(april, monthStr):
		return 4
	case strings.Contains(may, monthStr):
		return 5
	case strings.Contains(june, monthStr):
		return 6
	case strings.Contains(july, monthStr):
		return 7
	case strings.Contains(august, monthStr):
		return 8
	case strings.Contains(september, monthStr):
		return 9
	case strings.Contains(october, monthStr):
		return 10
	case strings.Contains(november, monthStr):
		return 11
	case strings.Contains(december, monthStr):
		return 12
	}
	return 0
}
