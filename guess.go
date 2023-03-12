package toss

import (
	"bytes"
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

func (o *Toss) guess() Schema {
	buf := new(bytes.Buffer)
	for _, record := range o.records {
		buf.WriteByte(record.Byte())
	}

	// P{n1}N{n2}模式
	ss := buf.String()
	pMax, nMax := size(ss, SchemaPositive), size(ss, SchemaNegative)

	for p := pMax; p >= 0; p-- {
		for n := nMax; n >= 0; n-- {
			var expr string
			var n0 string

			// 跳过
			if p+n == 0 {
				continue
			}

			xn := repetitions(p, n)
			if strings.HasPrefix(ss, SchemaPositive.String()) {
				n0 = fmt.Sprintf("%s%s", strings.Repeat("P", p), strings.Repeat("N", n))
				expr = fmt.Sprintf("^(P{%d}N{%d}){%d,}", p, n, xn)
			} else {
				n0 = fmt.Sprintf("%s%s", strings.Repeat("N", n), strings.Repeat("P", p))
				expr = fmt.Sprintf("^(N{%d}P{%d}){%d,}", n, p, xn)
			}

			s0 := regexp.MustCompile(expr).FindString(ss)
			if len(s0) > 0 {
				fmt.Printf("Pattern of %q matched %q: Prefer more than %d, and %d found.\n", n0, s0, xn, len(s0)/len(n0))

				schema := SchemaPositive
				if strings.HasPrefix(ss, SchemaNegative.String()) {
					schema = SchemaNegative
				}

				if p == 0 || n == 0 {
					return schema.Not()
				}

				return schema
			}
		}
	}

	return SchemaInvalid
}
