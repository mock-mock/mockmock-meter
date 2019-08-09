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
	// 時間計測：開始
	beforeTime := time.Now()
	log.Print("beforeTime：", beforeTime)

	// コマンドが違ったらバグ(設定ミス)なので、リターンする
	if !strings.Contains(req.Command, "end") {
		res := domain.SlackResponse{
			Text:         "BUG:command is not matched",
			Channel:      req.ChannelName,
			ResponseType: "in_channel",
		}
		return res
	}

	db, err := gorm.Open("postgres", utils.GetDBInfo())
	defer db.Close()
	LoggingPanic(err)

	// UserテーブルをPreloadして、MockMockレコードを持ってくる
	// update対象のUserIdを取得する。かつ、end_dateが埋まっていないもくもくレコードを探す。
	user := domain.User{
		SlackID: req.UserID,
	}
	db.Where("slack_id = ?", req.UserID).Preload("Mockmocks", "end_date = '0001-01-01 00:00:00'").Find(&user)
	// なければリターン。
	if len(user.Mockmocks) == 0 {
		res := domain.SlackResponse{
			Text:         "開始していません！新たにスタートするには、「/mock_start」してください。",
			Channel:      req.ChannelName,
			ResponseType: "in_channel",
		}
		return res
	} else if len(user.Mockmocks) > 1 {
		UnexpectedMockmocksLengthErr := errors.New("もくもく中のレコードが複数あります：" +
			strconv.Itoa(len(user.Mockmocks)))
		// 下のResponseを返したいので、Panicを起こさずエラーログを残す。Papertrailで拾う。
		// TODO：エラーだしユーザーはどうしようもないから、もくもくレコードを削除する仕様にしてもいいかも
		log.Print("error: " + UnexpectedMockmocksLengthErr.Error())
		res := domain.SlackResponse{
			Text:         "もくもく中のレコードが複数あります。Slack管理者へお伝えください。",
			Channel:      req.ChannelName,
			ResponseType: "in_channel",
		}
		return res
	}

	// もくもく中なのでそれのend_dateをアップデート
	endDate := time.Now()
	db.Model(&user.Mockmocks[0]).Update("end_date", endDate)

	// end_dateとstart_dateとの差を出してリターン。
	startDate := user.Mockmocks[0].StartDate
	duration := endDate.Sub(startDate)
	hour := duration.Truncate(time.Hour).Hours()               // Truncateは切り捨て
	min := duration.Round(time.Minute).Minutes() - (60 * hour) // Roundは四捨五入
	log.Print("hour: ", strconv.FormatFloat(hour, 'G', 4, 64))
	log.Print("min: ", strconv.FormatFloat(min, 'G', 4, 64))

	// 時間計測：終了
	afterTime := time.Now()
	log.Print("afterTime - beforeTime：", afterTime.Sub(beforeTime))
	log.Print("afterTime：", afterTime)

	res := domain.SlackResponse{
		Text: "もくもく終了！タイムは" + strconv.FormatFloat(hour, 'G', 4, 64) +
			"時間" + strconv.FormatFloat(min, 'G', 4, 64) + "分でした。お疲れ様でした！",
		Channel:      req.ChannelName,
		ResponseType: "in_channel",
	}
	return res
}
