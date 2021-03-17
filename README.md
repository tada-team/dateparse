[![Codecov coveragei build](https://github.com/tada-team/dateparse/actions/workflows/codecov.yml/badge.svg)](https://github.com/tada-team/dateparse/actions/workflows/codecov.yml) [![codecov](https://codecov.io/gh/tada-team/dateparse/branch/master/graph/badge.svg)](https://codecov.io/gh/tada-team/dateparse)
# dateparse

Лёгкий способ превратить пользовательский ввод даты во что-то машинопонятное.

Пользователи такие затейники, чего только не вводят, но мы пытаемся всё понять и простить:

```go
package main 

import (
    "time"
    "github.com/tada-team/dateparse"
)

func main() {
    date, message := dateparse.Parse("в следующий понедельник утром посмотреть код", nil)
    if date.IsZero() {
        panic("invalid date")
    }
    print("at:", date)
    print("do:", message)

    loc, err := time.LoadLocation("Europe/Moscow")
    if err != nil {
        panic(err)
    } 
    date, _ = dateparse.Parse("завтра", &dateparse.Opts{
        TodayEndHour: 20,
        Now:          time.Now().In(loc),
    })
    print(date)
}
```
