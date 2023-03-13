package toss

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

func rowsFn(name string) []int {
	buf, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("ReadFile() failure : %s", err.Error())
	}

	type Row struct {
		Result string `json:"lresult"`
	}

	var res []Row

	if err := json.Unmarshal(buf, &res); err != nil {
		log.Fatalf("Unmarshal() failure : %s", err.Error())
	}

	rows := make([]int, 0, len(res))
	for _, row := range res {
		i, err := strconv.Atoi(row.Result)
		if err != nil {
			log.Fatalf("Atoi() failure : %s", err.Error())
		}

		rows = append(rows, i)
	}

	return rows
}

func TestToss_Guess(t *testing.T) {
	rows := rowsFn("samples.json")
	fmt.Printf("历史数据共计 %d 条\n", len(rows))

	// 单双
	//pattern := oddEven

	// 大小
	//pattern := func(i int) Schema {
	//	if i >= 14 {
	//		return SchemaPositive
	//	}
	//
	//	return SchemaNegative
	//}

	// 尾数
	pattern := func(i int) Schema {
		if i%10 >= 5 {
			return SchemaPositive
		}

		return SchemaNegative
	}

	toss := NewToss(rows, pattern)

	// 开始猜测检查
	ns := rowsFn("samples_next.json")

	var dn, ok int
	for i := len(ns) - 1; i >= 0; i-- {
		log.Printf("开始 [第%d次] 猜测 ...... \n%s\n", len(ns)-i, toss)
		schema := toss.Guess()

		// 猜测结果
		toss.Add(ns[i])
		if schema == SchemaInvalid {
			continue
		}

		dn++
		result := pattern(ns[i])
		if result == schema {
			ok++
		}
	}

	log.Printf("猜测结果：测试样本量[ %d ] 猜测次数[ %d ] 猜测正确[ %d ] \n", len(ns), dn, ok)
	log.Printf("猜测结算：正确率[ %.2f%% ] 收益率[ %.2f ]\n", float64(ok*100)/float64(dn), float64(2*ok-dn)/float64(dn))
}
