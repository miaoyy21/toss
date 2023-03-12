package toss

import (
	"regexp"
	"testing"
)

type Patterna struct {
	name string
	expr string

	example string
}

var patterns = []Patterna{
	{
		name:    "以AB交替循环出现的次数至少大于4次",
		expr:    "^(AB){4,}",
		example: "ABABABAB",
	},
	{
		name:    "以ABB交替循环出现的次数至少大于4次",
		expr:    "^(AB{2}){4,}",
		example: "ABBABBABBAB",
	},
}

func TestFunc0(t *testing.T) {

	for _, pattern := range patterns {
		isMatched, err := regexp.MatchString(pattern.expr, pattern.example)
		if err != nil {
			t.Errorf("【%s】: %s\n", pattern.name, err.Error())
		}

		t.Logf("【%s】: Regexp %q Matched %q is %t\n", pattern.name, pattern.expr, pattern.example, isMatched)
	}

	rows := []int{14, 13, 3, 13, 23, 22, 14, 16, 17, 13, 22, 11, 9, 6, 4, 15, 3, 4, 15, 15, 16}
	toss := NewToss(rows, func(i int) Schema {
		if i%2 == 1 {
			return SchemaPositive
		}

		return SchemaNegative
	})

	toss.Add(15)

	t.Log(toss)
}
