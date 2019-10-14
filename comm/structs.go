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

const HOLIDAY_BACK_TYPE_JSON = "json"
const HOLIDAY_BACK_TYPE_TEXT = "text"

const (
	HOLIDAY_TYPE_WORK     = 0 //工作日
	HOLIDAY_TYPE_HOLIDAY  = 1 //假日
	HOLIDAY_TYPE_FESTIVAL = 2 //节日
)

//{
//	"status": 1,
//	"type": 2,
//	"typename": "节日",
//	"day": "20191001",
//	"unixtime": 1569859200,
//	"yearname": "己亥",
//	"nonglicn": "九月初三",
//	"nongli": "9-3",
//	"shengxiao": "猪",
//	"jieqi": "秋分后",
//	"weekcn": "二",
//	"week1": "Tue",
//	"week2": "2",
//	"week3": "Tuesday",
//	"daynum": "273",
//	"weeknum": "40",
//	"avoid": "入殓.安葬.开市.交易",
//	"suit": "祭祀.诸事不宜"
//}
type FullHoliday struct {
	Status    int    `json:"status"`
	Type      int    `json:"type"`
	Typename  string `json:"typename"`
	Day       string `json:"day"`
	Unixtime  int64  `json:"unixtime"`
	Yearname  string `json:"yearname"`
	Nonglicn  string `json:"nonglicn"`
	Nongli    string `json:"nongli"`
	Shengxiao string `json:"shengxiao"`
	Jieqi     string `json:"jieqi"`
	Weekcn    string `json:"weekcn"`
	Week1     string `json:"week1"`
	Week2     string `json:"week2"`
	Week3     string `json:"week3"`
	Daynum    string `json:"daynum"`
	Weeknum   string `json:"weeknum"`
	Avoid     string `json:"avoid"`
	Suit      string `json:"suit"`
}
