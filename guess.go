package toss

import (
	"bytes"
	"fmt"
	"regexp"
)

func size(s string, schema Schema) (min int, max int) {
	re := regexp.MustCompile(fmt.Sprintf("(%s)+", schema))

	ss := re.FindAllString(s, -1)
	for _, s0 := range ss {
		if len(s0) < min {
			min = len(s0)
		}

		if len(s0) > max {
			max = len(s0)
		}
	}

	return
}

func (o *Toss) guess() Schema {
	buf := new(bytes.Buffer)
	for _, record := range o.records {
		buf.WriteByte(record.Byte())
	}

	//获取最小
	pn, pm := size(buf.String(), SchemaPositive)
	nn, nm := size(buf.String(), SchemaNegative)

	fmt.Printf("P [%d,%d]\n", pn, pm)
	fmt.Printf("N [%d,%d]\n", nn, nm)

	return SchemaNegative
}
