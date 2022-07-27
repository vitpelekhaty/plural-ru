# plural-ru

Выбор формы существительного (множественное или единственное число, именительный или родительный падеж), связанного с числительным, для Go

## Использование

```go
package main

import (
	"fmt"
	
	ru "github.com/vitpelekhaty/plural-ru"
)

func main() {
	msgCount := 245
	
	str := fmt.Sprintf("получено %d %s", msgCount, ru.Plural(msgCount, "сообщение", "сообщения", "сообщений"))
	
	fmt.Println(str)
}

// Вывод:
//   получено 245 сообщений
```
