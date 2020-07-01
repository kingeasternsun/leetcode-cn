package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type DidStatusReqInfoStu struct {
	Mac          string    `json:"did"`
	Type         string    `json:"type"`
	OrderId      string    `json:"orderid"`
	ProcessName  string    `json:"processname"`
	ProcessNum   int       `json:"processcount"`
	ProcessId    int       `json:"processindex"`
	Result       int       `json:"result"`
	Info         string    `json:"info"`
	ReportTime   string    `json:"reporttime"`
	ReportTm     time.Time `json:"reporttm"` //这个仅仅用于获取表名
	OemFactoryId string    `json:"oemfactoryid"`
	UserId       string    `json:"userid"`
	Md5          string    `json:"md5"`
	VerTag       string    `json:"vertag"`
	//以下是为了校验烧录的文件是否正确
	CheckSumType   string `json:"checksumtype"` //对应脚本中的文件类型 manu_name,cal_name
	CheckSumFile   string `json:"checksumfile"` //文件名
	CheckSum       string `json:"checksum"`     //md5
	MacStatusTable string `json:"macstatustable"`
}

func main() {
	a := DidStatusReqInfoStu{}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// dp[i][j][0]

func maxProfit(prices []int) int {

	maxK := 3
	pLen := len(prices)
	if pLen <= 1 {
		return 0
	}

	dp := make([][][2]int, pLen)
	for i := 0; i < pLen; i++ {
		dp[i] = make([][2]int, maxK)
	}
	for i := 0; i < pLen; i++ {

		for k := 1; k < maxK; k++ {

			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
			} else {
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
				dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
			}
		}

	}

	return dp[pLen-1][maxK-1][0]
}
