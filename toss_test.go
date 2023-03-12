package toss

import (
	"testing"
)

type testTossData struct {
	rows    []int
	pattern func(i int) Schema

	row    int
	expect Schema
}

var testTossMap = []testTossData{
	{
		rows:    []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1},
		pattern: oddEven,

		row:    1,
		expect: SchemaPositive,
	},
	{
		rows:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		pattern: oddEven,

		row:    0,
		expect: SchemaPositive,
	},
	{
		rows:    []int{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1},
		pattern: oddEven,

		row:    0,
		expect: SchemaNegative,
	},
}

func oddEven(i int) Schema {
	if i%2 == 1 {
		return SchemaPositive
	}

	return SchemaNegative
}

func TestToss(t *testing.T) {
	for i, xs := range testTossMap {
		toss := NewToss(xs.rows, xs.pattern)
		toss.Add(xs.row)

		t.Logf("%s", toss)
		schema := toss.Guess()
		if schema == xs.expect {
			t.Logf("[%02d] Guess %q Successful ...", i+1, schema)
		} else {
			t.Fatalf("[%02d] Guess NOT PASSED, want %q but got %q", i+1, xs.expect, schema)
		}
	}
}
