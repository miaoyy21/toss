package xxx

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

var issue string
var bets = make(map[string]struct{})

func Run() error {
	id, token := "31591499", "cbj7s576p3se6c87194kwqo1c1w2cq87sau8lc2s"
	unix, code := "1680178143", "a6748dba269e72b5ea7bb9bb7c4ee619"
	device := "0E6EE3CC-8184-4CD7-B163-50AE8AD4516F"
	decision := 1.25
	power := 50 // 投注倍率：目标中奖金额为投注倍率*1000

	// 查询近期历史
	hisRequest := QHistoryRequest{
		PageSize: 1000,
		PType:    "3",
		Unix:     unix,
		KeyCode:  code,
		DeviceId: device,
		UserId:   id,
		Token:    token,
	}

	var hisResponse QHistory

	// 执行查询开奖历史
	err := execute("GET", "http://manorapp.pceggs.com/IFS/Manor28/Manor28_Analyse_History.ashx", hisRequest, &hisResponse)
	if err != nil {
		return fmt.Errorf("查询开奖历史出现错误：%s", err.Error())
	}

	// 开奖历史是否存在错误
	if hisResponse.Status != 0 {
		return fmt.Errorf("查询开奖历史成功，但存在错误返回：(%d) %s", hisResponse.Status, hisResponse.Msg)
	}

	// 开奖历史为空
	if len(hisResponse.Data.Items) < 1 {
		return errors.New("没有查询到开奖历史")
	}

	// 最新开奖期数
	nowIssue := hisResponse.Data.Items[0].Issue
	if strings.EqualFold(nowIssue, issue) {
		log.Printf("本期开奖期数【%s】，还没到开奖时间，等待下次执行 ...\n", nowIssue)
		return nil
	}

	// 开奖结果
	res := hisResponse.Data.Items[0].Result
	if len(bets) == 0 {
		log.Printf("本期开奖期数【%s】，开奖结果【%s】 ...\n", nowIssue, res)
	} else {
		if _, ok := bets[res]; ok {
			log.Printf("本期开奖期数【%s】，开奖结果【%s】，已中奖 [✓]...\n", nowIssue, res)
		} else {
			log.Printf("本期开奖期数【%s】，开奖结果【%s】，没有中奖 [×]...\n", nowIssue, res)
		}
	}
	issue = nowIssue

	spaces := make(map[int]int)
	for index, item := range hisResponse.Data.Items {
		result, err := strconv.Atoi(item.Result)
		if err != nil {
			return err
		}

		// Exists
		if _, ok := spaces[result]; ok {
			continue
		}

		spaces[result] = index + 1
	}

	iNextIssue, err := strconv.Atoi(nowIssue)
	if err != nil {
		return err
	}
	nextIssue := strconv.Itoa(iNextIssue + 1)

	// 开奖较频繁的结果，如果大于2/3，那么再进行一次或两次退化
	target, price := make([]int, 0), 1000.0
	target, price = getTarget(spaces, decision)
	if price >= 667 {
		log.Printf("下期开奖期数【%s】，预测中奖率【%.2f%%】，第一次退化 ...\n", nextIssue, price/10)
		target, price = getTarget(spaces, decision*0.9)
		if price >= 667 {
			log.Printf("下期开奖期数【%s】，预测中奖率【%.2f%%】，第二次退化 ...\n", nextIssue, price/10)
			target, price = getTarget(spaces, decision*0.9*0.9)
		}
	}
	sort.Ints(target)

	log.Printf("下期开奖期数【%s】，预测中奖率【%.2f%%】，即将投注 %v ...\n", nextIssue, price/10, target)

	var total int
	bets = make(map[string]struct{}) // 清空前一次投注结果
	for _, result := range target {
		gold := int(float64(1000*power) / float64(standard[result]))

		betRequest := XBetRequest{
			Issue:    nextIssue,
			GoldEggs: gold,
			Numbers:  result,
			Unix:     unix,
			Keycode:  code,
			PType:    "3",
			DeviceId: device,
			Userid:   id,
			Token:    token,
		}

		var betResponse XBet
		err := execute("GET", "http://manorapp.pceggs.com/IFS/Manor28/Manor28_Betting_1.ashx", betRequest, &betResponse)
		if err != nil {
			return fmt.Errorf("下期开奖期数【%s】，执行押注[%d]，出现错误：%s", nextIssue, result, err.Error())
		}

		if betResponse.Status != 0 {
			return fmt.Errorf("下期开奖期数【%s】，执行押注[%d]，服务器返回错误信息：%s", nextIssue, result, betResponse.Msg)
		}

		total = total + gold
		bets[strconv.Itoa(result)] = struct{}{}
	}
	log.Printf("下期开奖期数【%s】，投入 %d，押注成功 >>>>>>>>>> \n", nextIssue, total)

	return nil
}

func getTarget(spaces map[int]int, decision float64) ([]int, float64) {
	var price float64
	target := make([]int, 0)
	for result, space := range spaces {
		if int(decision*float64(standard[result])) > space {
			target = append(target, result)
			price = price + float64(1000)/float64(standard[result])
		}
	}

	return target, price
}
