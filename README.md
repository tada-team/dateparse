# dateparse

Лёгкий способ превратить пользовательский ввод даты во что-то машинопонятное.

Пользователи такие затейники, чего только не вводят, но мы пытаемся всё понять и простить:

```go
package main 

import "github.com/tada-team/dateparse"

func main() {
    date, message := dateparse.Parse("в следующий понедельник утром посмотреть код", nil)
    print("at:", date)
    print("do:", message)
}
```
