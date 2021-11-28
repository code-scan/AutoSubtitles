package pkg

import (
	"fmt"
)

type Aliyun struct {
	AK     string
	AS     string
	AppKey string
}
type ZimuResult struct {
	TaskID      string `json:"TaskId"`
	RequestID   string `json:"RequestId"`
	StatusText  string `json:"StatusText"`
	BizDuration int    `json:"BizDuration"`
	SolveTime   int64  `json:"SolveTime"`
	StatusCode  int    `json:"StatusCode"`
	Result      Result `json:"Result"`
}
type Sentences struct {
	EndTime         int64   `json:"EndTime"`
	SilenceDuration int     `json:"SilenceDuration"`
	BeginTime       int64   `json:"BeginTime"`
	Text            string  `json:"Text"`
	ChannelID       int     `json:"ChannelId"`
	SpeechRate      int     `json:"SpeechRate"`
	EmotionValue    float64 `json:"EmotionValue"`
}
type Result struct {
	Sentences []Sentences `json:"Sentences"`
}

func Int2time(t int64) string {
	var (
		second, min, hours, millisecond int64
	)
	millisecond = t % 1000
	second = t / 1000
	if second > 59 {
		min = (t / 1000) / 60
		second = second % 60
	}
	if min > 59 {
		hours = (t / 1000) / 3600
		min = min % 60
	}
	return fmt.Sprintf("%d:%d:%d,%d", hours, min, second, millisecond)

}
