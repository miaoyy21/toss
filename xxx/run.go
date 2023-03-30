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

func Run() error {
	id, token := "31591499", "cbj7s576p3se6c87194kwqo1c1w2cq87sau8lc2s"
	unix, code := "1680178143", "a6748dba269e72b5ea7bb9bb7c4ee619"
	device := "0E6EE3CC-8184-4CD7-B163-50AE8AD4516F"
	decision := 0.75

	log.Printf("开始执行查询历史记录 ...\n")

	// 查询近期历史
	hisRequest := QHistoryRequest{
		PageSize: 200,
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
		log.Printf("最新开奖期数%q,还没到开奖时间，等待下次执行 ...\n", nowIssue)
		return nil
	}

	issue = nowIssue

	spaces, target := make(map[int]int), make([]int, 0)
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

	// 开奖较频繁的结果
	var price float64
	for result, space := range spaces {
		if int(decision*float64(standard[result])) > space {
			target = append(target, result)
			price = price + float64(1000)/float64(standard[result])
		}
	}
	sort.Ints(target)

	iNextIssue, err := strconv.Atoi(nowIssue)
	if err != nil {
		return err
	}

	nextIssue := strconv.Itoa(iNextIssue + 1)
	log.Printf("开奖期数%q,即将投注 %#v，覆盖率 %.2f%%...\n", nextIssue, target, price/10)

	var total int
	for _, result := range target {
		gold := int(float64(10000) / float64(standard[result]))

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
			return fmt.Errorf("执行押注[%d]，出现错误：%s", result, err.Error())
		}

		if betResponse.Status != 0 {
			return fmt.Errorf("执行押注[%d]，服务器返回错误信息：%s", result, err.Error())
		}

		total = total + gold
	}
	log.Printf("执行押注成功，累计投入 %d >>>>>>>>>> \n", total)

	return nil
}
