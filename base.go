package dateparse

var (
	dd   = `(30|31|0[1-9]|[1-2]\d|[1-9])`
	mm   = `(1[012]|0[1-9])`
	yy   = `(\d\d)`
	yyyy = `(20\d\dк?г?)`
	HH   = `(0[0-9]|1[0-9]|2[0-3]|[0-9])`
	MM   = `(0[0-9]|[0-5][0-9])`
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
