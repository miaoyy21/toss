package toss

import (
	"fmt"
	"regexp"
	"strings"
)

func size(s string, schema Schema) int {
	ss := regexp.MustCompile(fmt.Sprintf("(%s)+", schema)).FindAllString(s, -1)

	max := 0
	for _, s0 := range ss {
		if len(s0) > max {
			max = len(s0)
		}
	}

	return max
}

// neg: 是否开启反向竞猜
func guess(s string, neg bool) Schema {
	pMax, nMax := size(s, SchemaPositive), size(s, SchemaNegative)

	for p := pMax; p >= 0; p-- {
		for n := nMax; n >= 0; n-- {
			var expr string
			var n0 string

			// 跳过
			if p+n == 0 {
				continue
			}

			xn := repetitions(p, n)
			if strings.HasPrefix(s, SchemaPositive.String()) {
				n0 = fmt.Sprintf("%s%s", strings.Repeat("P", p), strings.Repeat("N", n))
				expr = fmt.Sprintf("^(P{%d}N{%d}){%d,}", p, n, xn)
			} else {
				n0 = fmt.Sprintf("%s%s", strings.Repeat("N", n), strings.Repeat("P", p))
				expr = fmt.Sprintf("^(N{%d}P{%d}){%d,}", n, p, xn)
			}

			s0 := regexp.MustCompile(expr).FindString(s)
			if len(s0) > 0 {
				fmt.Printf("Pattern of %q matched %q: Prefer more than %d, and %d found.\n", n0, s0, xn, len(s0)/len(n0))

				schema := SchemaPositive
				if strings.HasPrefix(s, SchemaNegative.String()) {
					schema = SchemaNegative
				}

				if p == 0 || n == 0 {
					return schema.Not()
				}

				return schema
			}
		}
	}

	// 反向竞猜
	if neg {
		fmt.Printf("Negation %q Guess.\n", SchemaPositive)
		s1 := fmt.Sprintf("%s%s", SchemaPositive, s)
		if guess(s1, false) != SchemaInvalid {
			return SchemaNegative
		}

		fmt.Printf("Negation %q Guess.\n", SchemaNegative)
		s2 := fmt.Sprintf("%s%s", SchemaNegative, s)
		if guess(s2, false) != SchemaInvalid {
			return SchemaPositive
		}
	}

	return SchemaInvalid
}
