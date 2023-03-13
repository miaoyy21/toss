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

	// P{p}N{n}: 其中 p + n !=0
	for p := pMax; p >= 0; p-- {
		for n := nMax; n >= 0; n-- {
			var expr string
			var n0 string

			// 跳过
			if p+n == 0 {
				continue
			}

			xn := repetitions(p, n)
			expr = fmt.Sprintf("^(%s{%d}%s{%d}){%d,}", SchemaPositive, p, SchemaNegative, n, xn)

			if strings.HasPrefix(s, SchemaPositive.String()) {
				n0 = fmt.Sprintf("%s%s", strings.Repeat(string(SchemaPositive), p), strings.Repeat(string(SchemaNegative), n))
				expr = fmt.Sprintf("^(%s{%d}%s{%d}){%d,}", SchemaPositive, p, SchemaNegative, n, xn)
			} else {
				n0 = fmt.Sprintf("%s%s", strings.Repeat(string(SchemaNegative), n), strings.Repeat(string(SchemaPositive), p))
				expr = fmt.Sprintf("^(%s{%d}%s{%d}){%d,}", SchemaNegative, n, SchemaPositive, p, xn)
			}

			s0 := regexp.MustCompile(expr).FindString(s)
			if len(s0) > 0 {
				fmt.Printf("Pattern of %q matched %q: Prefer more than %d, and %d found.\n", n0, s0, xn, len(s0)/len(n0))

				schema := SchemaPositive
				if strings.HasPrefix(s, SchemaNegative.String()) {
					schema = SchemaNegative
				}

				if p == 0 || n == 0 {
					return schema.Reverse()
				}

				return schema
			}
		}
	}

	if !neg {
		var expr string

		// P{p}N{n}: 其中 p >= 2 && n >= 2 && p1 != p2... && n1 != n2...
		expect := 3
		expr = fmt.Sprintf("^(%s{2,}%s{2,}){%d,}", Schema(s[0]), Schema(s[0]).Reverse(), expect)
		s0 := regexp.MustCompile(expr).FindString(s)
		if len(s0) > 0 {
			if s0[0] == s0[1] {
				n0 := strings.Count(s0, fmt.Sprintf("%s%s", Schema(s[0]), Schema(s[0]).Reverse()))
				fmt.Printf("Pattern of %s{2+}%s{2+} matched %q: Prefer more than %d, and %d found.\n", Schema(s[0]), Schema(s[0]).Reverse(), s0, expect, n0)
				return Schema(s[0]).Reverse()
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
