package dateparse

var (
	dayDD    = `(30|31|0[1-9]|[1-2]\d|[1-9])`
	monthMM  = `(1[012]|0[1-9])`
	yearYY   = `(\d\d)`
	yearYYYY = `(20\d\dк?г?)`
	hourHH   = `(0[0-9]|1[0-9]|2[0-3]|[0-9])`
	minuteMM = `(0[0-9]|[0-5][0-9])`
)

var (
	morning  = `утром|утро|утра|morning|a.m`
	evening  = `вечером|вечер|вечера|evening|p.m`
	midnight = "полночь|ночью|midnight"
	noon     = "днем|полдень|noon|midday"
)

var (
	today              = `сегодня|today`
	tomorrow           = `завтра|tomorrow`
	afterTomorrow      = `послезавтра|after tomorrow|aftertomorrow`
	afterAfterTomorrow = `послепослезавтра|after after tomorrow|afteraftertomorrow`
	yesterday          = `вчера|yesterday`
)
