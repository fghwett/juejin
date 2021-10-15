package task

type Response struct {
	ErrNo  int         `json:"err_no"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

// 签到返回值
type SignResponse struct {
	IncrPoint int `json:"incr_point"`
	SumPoint  int `json:"sum_point"`
}

// 抽奖返回值
type LotteryResponse struct {
	Id           int    `json:"id"`
	LotteryId    string `json:"lottery_id"`
	LotteryName  string `json:"lottery_name"`
	LotteryType  int    `json:"lottery_type"`
	LotteryImage string `json:"lottery_image"`
	HistoryId    string `json:"history_id"`
}
