package command

import (
	"github.com/mock-mock/mockmock-meter/backend/domain"
	"github.com/mock-mock/mockmock-meter/backend/utils"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Start(req domain.SlackRequest) domain.SlackResponse {
	// 時間計測：開始
	beforeTime := time.Now()
	log.Print("beforeTime：", beforeTime)

	//コマンドが違ったらバグなので、リターンする
	if !strings.Contains(req.Command, "start") {
		res := domain.SlackResponse{
			Text:         "BUG:command is not matched",
			Channel:      req.ChannelName,
			ResponseType: "in_channel",
		}
		return res
	}

	// UserテーブルをPreloadして、MockMockレコードを持ってくる
	db, err := gorm.Open("postgres", utils.GetDBInfo())
	defer db.Close()
	LoggingPanic(err)

	// SlackIdからUserId取得＋現在のもくもくレコードを取得
	user := domain.User{
		SlackId: req.UserID,
	}
	db.Where("slack_id = ?", req.UserID).Preload("Mockmocks", "end_date = '0001-01-01 00:00:00'").Find(&user)
	log.Print(user)
	// Userレコードが取れたかチェック
	if user.ID == 0 {
		// userレコードを作る
		user.Name = req.UserName
		db.Create(&user)
	}

	// もくもく中かチェック
	if len(user.Mockmocks) != 0 {
		res := domain.SlackResponse{
			Text:         "まだもくもく中です！新たにスタートするには、「/mock_end」してください。",
			Channel:      req.ChannelName,
			ResponseType: "in_channel",
		}
		return res
	}

	// Insertする
	insertMock(db, user.ID)
	res := domain.SlackResponse{
		Text:         "もくもくスタート！",
		Channel:      req.ChannelName,
		ResponseType: "in_channel",
	}
	// 時間計測：終了
	afterTime := time.Now()
	log.Print("afterTime - beforeTime：", afterTime.Sub(beforeTime))
	log.Print("afterTime：", afterTime)
	return res
}

func insertMock(db *gorm.DB, userID int) {
	mock := domain.Mockmock{
		StartDate: time.Now(),
		UserID:    userID,
		Category:  "もくもく",
	}
	db.Create(&mock)
}

func LoggingPanic(err error) {
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[mockmock] ")
		log.Panic(err)
	}
	return
}
