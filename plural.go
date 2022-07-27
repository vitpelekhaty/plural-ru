package plural_ru

import (
	"strconv"
)

// Plural возвращает слово, связанное с числительным, в соответствующей форме (падеже и числе)
//
// Пример:
// 	 val := plural_ru.Plural(12, "сообщение", "сообщения", "сообщений")
//   fmt.Println(val)
//
func Plural(quantity int, nominative, accusative, plural string) string {
	sq := strconv.Itoa(quantity)

	var s2 string

	if len(sq) > 1 {
		s2 = sq[len(sq)-2:]
	}

	s1 := sq[len(sq)-1:]

	d2, _ := strconv.Atoi(s2)
	d1, _ := strconv.Atoi(s1)

	if d1 == 1 && !(d2 == 11) {
		return nominative
	} else if d1 > 1 && d1 < 5 && !(d2 > 11 && d2 < 15) {
		return accusative
	} else {
		return plural
	}
}
