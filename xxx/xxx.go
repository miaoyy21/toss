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
	power := 20 // 投注倍率：目标中奖金额为投注倍率*1000

	// 查询近期历史
	hisRequest := QHistoryRequest{
		PageSize: 500,
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
	target, price = getTarget(spaces)
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

func getTarget(spaces map[int]int) ([]int, float64) {
	type Space struct {
		Result int
		Space  int
	}

	newSpaces := make([]Space, 0, len(spaces))
	for result, space := range spaces {
		newSpaces = append(newSpaces, Space{Result: result, Space: space})
	}
	sort.Slice(newSpaces, func(i, j int) bool {
		return float64(newSpaces[i].Space)/float64(standard[newSpaces[i].Result]) < float64(newSpaces[j].Space)/float64(standard[newSpaces[j].Result])
	})

	target, price := make([]int, 0), 0.0
	for _, newSpace := range newSpaces {
		if price+float64(1000)/float64(standard[newSpace.Result]) >= 550 {
			break
		}

		price = price + float64(1000)/float64(standard[newSpace.Result])
		target = append(target, newSpace.Result)
	}

	return target, price
}
