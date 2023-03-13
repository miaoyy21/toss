package toss

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestToss_Guess(t *testing.T) {
	buf, err := os.ReadFile("samples.json")
	if err != nil {
		t.Fatalf("ReadFile() failure : %s", err.Error())
	}

	type Row struct {
		Result string `json:"lresult"`
	}

	var res []Row

	if err := json.Unmarshal(buf, &res); err != nil {
		t.Fatalf("Unmarshal() failure : %s", err.Error())
	}

	rows := make([]int, 0, len(res))
	for _, row := range res {
		i, err := strconv.Atoi(row.Result)
		if err != nil {
			t.Fatalf("Atoi() failure : %s", err.Error())
		}

		rows = append(rows, i)
	}

	fmt.Printf("Rows Count is %d\n", len(rows))
}
