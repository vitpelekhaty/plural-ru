package plural_ru

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlural(t *testing.T) {
	var cases = [...]struct {
		quantity   int
		nominative string
		accusative string
		plural     string
		expected   string
	}{
		{quantity: 1, nominative: "тест", accusative: "теста", plural: "тестов", expected: "тест"},
		{quantity: 10, nominative: "тест", accusative: "теста", plural: "тестов", expected: "тестов"},
		{quantity: 2, nominative: "тест", accusative: "теста", plural: "тестов", expected: "теста"},
		{quantity: 22, nominative: "тест", accusative: "теста", plural: "тестов", expected: "теста"},
		{quantity: 0, nominative: "тест", accusative: "теста", plural: "тестов", expected: "тестов"},
		{quantity: 1018, nominative: "тест", accusative: "теста", plural: "тестов", expected: "тестов"},
		{quantity: 7761, nominative: "тест", accusative: "теста", plural: "тестов", expected: "тест"},
		{quantity: 3252, nominative: "тест", accusative: "теста", plural: "тестов", expected: "теста"},
		{quantity: 8811, nominative: "тест", accusative: "теста", plural: "тестов", expected: "тестов"},
	}

	for _, test := range cases {
		v := Plural(test.quantity, test.nominative, test.accusative, test.plural)
		require.Equalf(t, test.expected, v, "%v", test)
	}
}
