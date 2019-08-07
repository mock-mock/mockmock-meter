package command

import (
	"github.com/mock-mock/mockmock-meter/backend/domain"
	"github.com/mock-mock/mockmock-meter/backend/utils"
	"log"
	"strings"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Start(req domain.SlackRequest) domain.SlackResponse {
	// memo: https://script.google.com/macros/s/AKfycbwBeiTRWprDV9RJWgH8AqHSgZh5QB-7EEgyYOMVoquDq-27GELg/exec
	log.Print("get req")
	bf_t := time.Now()
	//コマンドが違ったらバグなので、リターンする
	if !strings.Contains(req.Command, "start") {
		res := domain.SlackResponse{Text: "BUG:command is not matched", Channel: req.ChannelName, ResponseType: "in_channel"}
		return res 
	}

	// UserテーブルをPreloadして、MockMockレコードを持ってくる
	db, err := gorm.Open("postgres", utils.GetDBInfo())
	defer db.Close()
	loggingPanic(err)

	user := domain.User{
		SlackId: req.UserID,
	}
	// SlackIdからUserId取得＋現在のもくもくレコードを取得
	db.Where("slack_id = ?", req.UserID).Preload("Mockmocks", "end_date = '0001-01-01 00:00:00'").Find(&user)
	log.Print(user)
	// もくもく中かチェック
	if len(user.Mockmocks) != 0 {
		res := domain.SlackResponse{Text: "まだもくもく中です！新たにスタートするには、「/mock_end」してください。", Channel: req.ChannelName, ResponseType: "in_channel"}
		return res
	}

	// Insertする
	insertMock(db, user.ID)
	log.Print("inserted!!!")
	res := domain.SlackResponse{Text: "もくもくスタート！", Channel: req.ChannelName, ResponseType: "in_channel"}
	log.Print(res)
	af_t := time.Now()
	log.Print(af_t.Sub(bf_t))
	return res
}

func insertMock(db *gorm.DB, userId int) {
	mock := domain.Mockmock{
		StartDate: time.Now(),
		UserID: userId,
		Category: "もくもく",
	}
	db.Create(&mock)
}

func loggingPanic(err error) {
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[mockmock] ")
		log.Panic(err)
	}
	return
}
