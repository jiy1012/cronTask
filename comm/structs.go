package comm

//钉钉机器人报警结构体
type TextStruct struct {
	Content string `json:"content"`
}
type AtStruct struct {
	Atmobiles []string `json:"atMobiles"`
	Isatall   bool     `json:"isAtAll"`
}
type DDRobotStruct struct {
	Msgtype string     `json:"msgtype"`
	Text    TextStruct `json:"text"`
	At      AtStruct   `json:"at"`
}
