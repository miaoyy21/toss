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
	fmt.Printf("样本的历史数据共计 %d 条\n", len(rows))

	// 单双
	p1 := oddEven

	// 大小
	p2 := func(i int) Schema {
		if i >= 14 {
			return SchemaPositive
		}

		return SchemaNegative
	}

	// 尾数
	p3 := func(i int) Schema {
		if i%10 >= 5 {
			return SchemaPositive
		}

		return SchemaNegative
	}

	// 靠中
	p4 := func(i int) Schema {
		if i >= 11 && i <= 16 {
			return SchemaPositive
		}

		return SchemaNegative
	}

	// 自定义
	p5 := func(i int) Schema {
		if i == 1 || i == 2 || i == 3 || i == 5 || i == 8 ||
			i == 10 || i == 12 || i == 14 || i == 17 || i == 19 ||
			i == 20 || i == 21 || i == 23 || i == 24 || i == 25 || i == 27 {
			return SchemaPositive
		}

		return SchemaNegative
	}

	var dn, ok, nx0, nxx int

	capacity := 100
	toss := NewToss(rows[len(rows)-capacity:], p1)
	ns := rows[:len(rows)-capacity]

	for i := len(ns) - 1; i >= 0; i-- {
		toss.ResetPattern(p1)
		log.Printf("开始 [第%d次] 1.单双模式 猜测 ...... \n%s\n", len(ns)-i, toss)
		schema := toss.Guess()

		var result Schema

		// 猜测结果
		if schema == SchemaInvalid {
			toss.ResetPattern(p2)
			log.Printf("开始 [第%d次] 2.大小模式 猜测 ...... \n%s\n", len(ns)-i, toss)

			schema = toss.Guess()
			if schema == SchemaInvalid {
				toss.ResetPattern(p3)
				log.Printf("开始 [第%d次] 3.大小尾数模式 猜测 ...... \n%s\n", len(ns)-i, toss)

				schema = toss.Guess()
				if schema == SchemaInvalid {
					toss.ResetPattern(p4)
					log.Printf("开始 [第%d次] 4.靠中模式 猜测 ...... \n%s\n", len(ns)-i, toss)

					schema = toss.Guess()
					if schema == SchemaInvalid {
						toss.ResetPattern(p5)
						log.Printf("开始 [第%d次] 5.自定义模式 猜测 ...... \n%s\n", len(ns)-i, toss)

						schema = toss.Guess()
						if schema == SchemaInvalid {
							toss.Add(ns[i])
							continue
						} else {
							result = p5(ns[i])
						}
					} else {
						result = p4(ns[i])
					}
				} else {
					result = p3(ns[i])
				}
			} else {
				result = p2(ns[i])
			}
		} else {
			result = p1(ns[i])
		}

		dn++
		if result == schema {
			ok++
			log.Printf("开始 [第%d次] 猜测正确 [✓]\n", len(ns)-i)

			if nx0 > nxx {
				nxx = nx0
			}
			nx0 = 0
		} else {
			nx0++
			log.Printf("开始 [第%d次] 猜测错误 [×]\n", len(ns)-i)
		}

		toss.Add(ns[i])
	}

	log.Printf("猜测结果：测试样本量[ %d ] 猜测次数[ %d ] 猜测正确[ %d ] \n", len(ns), dn, ok)
	log.Printf("猜测结算：正确率[ %.2f%% ] 最大连错次数[ %d ] 收益率[ %.2f ]\n", float64(ok*100)/float64(dn), nxx, float64(2*ok-dn)/float64(dn))
}
