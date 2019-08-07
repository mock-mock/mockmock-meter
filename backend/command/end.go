package command

import (
	"errors"
	"github.com/mock-mock/mockmock-meter/backend/domain"
	"github.com/mock-mock/mockmock-meter/backend/utils"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// End Mock-Mock
func End(req domain.SlackRequest) domain.SlackResponse {
	/**
	req sample
	1 DJ74B9YJV
	2 directmessage
	3 /mock_end
	4 https://hooks.slack.com/commands/TJ20YR2Q2/696233068646/wNAAbSugQdJXpb5cVgYF7SCX
	5 mock-mock
	6 TJ20YR2Q2
	7
	8 2CkUfxeyQUFjVhaKzq6rBWrv
	9 696233068678.614032852818.2100c9ab0db9e5f233328f37023637f2
	10 UJ74B9N1X
	11 oshikawatakuya
	*/
	// fmt.Println("1" + req.ChannelID)
	// fmt.Println("2" + req.ChannelName)
	// fmt.Println("3" + req.Command)
	// fmt.Println("4" + req.ResponseURL)
	// fmt.Println("5" + req.TeamDomain)
	// fmt.Println("6" + req.TeamID)
	// fmt.Println("7" + req.Text)
	// fmt.Println("8" + req.Token)
	// fmt.Println("9" + req.TriggerID)
	// fmt.Println("10" + req.UserID)
	// fmt.Println("11" + req.UserName)

	// 時間計測：開始
	beforeTime := time.Now()
	log.Print("beforeTime：", beforeTime)

	//コマンドが違ったらバグなので、リターンする
	if !strings.Contains(req.Command, "end") {
		res := domain.SlackResponse{Text: "BUG:command is not matched", Channel: req.ChannelName, ResponseType: "in_channel"}
		return res
	}

	// UserテーブルをPreloadして、MockMockレコードを持ってくる
	db, err := gorm.Open("postgres", utils.GetDBInfo())
	defer db.Close()
	LoggingPanic(err)

	//update対象のUserIdを取得する。かつ、end_dateが埋まっていないもくもくレコードを探す。
	user := domain.User{
		SlackId: req.UserID,
	}
	db.Where("slack_id = ?", req.UserID).Preload("Mockmocks", "end_date = '0001-01-01 00:00:00'").Find(&user)
	//なければリターン。
	if len(user.Mockmocks) == 0 {
		res := domain.SlackResponse{
			Text:         "開始していません！新たにスタートするには、「/mock_start」してください。",
			Channel:      req.ChannelName,
			ResponseType: "in_channel"}
		return res
	} else if len(user.Mockmocks) > 1 {
		UnexpectedMockmocksLengthErr := errors.New("もくもく中のレコードが複数あります：" +
			strconv.Itoa(len(user.Mockmocks)))
		LoggingPanic(UnexpectedMockmocksLengthErr)
		res := domain.SlackResponse{
			Text:         "もくもく中のレコードが複数あります",
			Channel:      req.ChannelName,
			ResponseType: "in_channel"}
		return res
	}

	// もくもく中なのでそれをアップデート
	endDate := time.Now()
	db.Model(&user.Mockmocks[0]).Update("end_date", endDate)

	// start_dateとの差を出してリターン。
	startDate := user.Mockmocks[0].StartDate
	duration := endDate.Sub(startDate)

	// 時間計測：終了
	afterTime := time.Now()
	log.Print("afterTime - beforeTime：", afterTime.Sub(beforeTime))
	log.Print("afterTime：", afterTime)

	res := domain.SlackResponse{
		Text:         "もくもく終了！タイムは" + duration.String() + "でした。お疲れ様でした！",
		Channel:      req.ChannelName,
		ResponseType: "in_channel",
	}
	return res
}
