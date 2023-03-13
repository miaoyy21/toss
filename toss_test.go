package toss

import (
	"fmt"
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
		expect: SchemaNegative,
	},
	{
		rows:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		pattern: oddEven,

		row:    0,
		expect: SchemaNegative,
	},
	{
		rows:    []int{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1},
		pattern: oddEven,

		row:    0,
		expect: SchemaPositive,
	},
	{
		rows:    []int{0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1},
		pattern: oddEven,

		row:    1,
		expect: SchemaNegative,
	},
	{
		rows:    []int{1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1},
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

		fmt.Printf("%s\n", toss)
		schema := toss.Guess()
		if schema == xs.expect {
			fmt.Printf("[%02d] Guess %q Successful. [✓]\n", i+1, schema)
		} else {
			t.Fatalf("[%02d] Guess NOT PASSED, want %q but got %q. [×]\n", i+1, xs.expect, schema)
		}
	}
}
