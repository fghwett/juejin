package task

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fghwett/juejin/config"
	"github.com/fghwett/juejin/util"
)

type Task struct {
	cookie    string
	userAgent string
	client    *http.Client
	result    []string
}

func New(config *config.Config) *Task {
	return &Task{
		cookie:    config.Cookie,
		userAgent: config.UserAgent,
		client:    &http.Client{},
		result:    []string{"==== 掘金保活任务 ===="},
	}
}

func (t *Task) Do() {
	if err := t.signTask(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【签到任务】：失败 %s", err))
		return
	}

	util.SmallSleep(1000, 3000)

	if err := t.lotteryTask(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【抽奖任务】：失败 %s", err))
		return
	}

	util.SmallSleep(3000, 5000)

	if err := t.getPointTask(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【积分查询】：失败 %s", err))
		return
	}
}

func (t *Task) signTask() error {
	reqUrl := "https://api.juejin.cn/growth_api/v1/check_in"
	req, err := http.NewRequest(http.MethodPost, reqUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("origin", "https://juejin.cn")
	req.Header.Set("referer", "https://juejin.cn/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", t.userAgent)
	req.Header.Set("cookie", t.cookie)

	resp, err := t.client.Do(req)

	response := &Response{Data: &SignResponse{}}
	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	signResp := response.Data.(*SignResponse)
	t.result = append(t.result, fmt.Sprintf("【签到任务】：成功 获得积分%d 累计积分%d", signResp.IncrPoint, signResp.SumPoint))

	return nil
}

func (t *Task) lotteryTask() error {
	reqUrl := "https://api.juejin.cn/growth_api/v1/lottery/draw"
	req, err := http.NewRequest(http.MethodPost, reqUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("origin", "https://juejin.cn")
	req.Header.Set("referer", "https://juejin.cn/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", t.userAgent)
	req.Header.Set("cookie", t.cookie)

	resp, err := t.client.Do(req)

	response := &Response{Data: &LotteryResponse{}}
	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	lotteryResp := response.Data.(*LotteryResponse)
	t.result = append(t.result, fmt.Sprintf("【抽奖任务】：成功 获得%s", lotteryResp.LotteryName))

	return nil
}

func (t *Task) getPointTask() error {
	reqUrl := "https://api.juejin.cn/growth_api/v1/get_cur_point"
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("origin", "https://juejin.cn")
	req.Header.Set("referer", "https://juejin.cn/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", t.userAgent)
	req.Header.Set("cookie", t.cookie)

	resp, err := t.client.Do(req)

	response := &Response{}
	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	t.result = append(t.result, fmt.Sprintf("【积分查询】：成功 余额%d", int(response.Data.(float64))))

	return nil
}

func (t *Task) GetResult() string {
	return strings.Join(t.result, " \n\n ")
}
