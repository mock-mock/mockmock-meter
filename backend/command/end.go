package command

import (
	"fmt"
	"strings"
	"github.com/mock-mock/mockmock-meter/backend/domain"
)

// End Mock-Mock
func End(req domain.SlackRequest) domain.SlackResponse {
	/** 
	sample
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
	fmt.Println("1" + req.ChannelID)
	fmt.Println("2" + req.ChannelName)
	fmt.Println("3" + req.Command)
	fmt.Println("4" + req.ResponseURL)
	fmt.Println("5" + req.TeamDomain)
	fmt.Println("6" + req.TeamID)
	fmt.Println("7" + req.Text)
	fmt.Println("8" + req.Token)
	fmt.Println("9" + req.TriggerID)
	fmt.Println("10" + req.UserID)
	fmt.Println("11" + req.UserName)

	//コマンドが違ったらバグなので、リターンする
	if !strings.Contains(req.Command, "end") {
		res := domain.SlackResponse{Text: "BUG:command is not matched", Channel: req.ChannelName, ResponseType: "in_channel"}
		return res 
	}

	//update対象のUserIdを取得する

	//取得したUserIdで、end_dateが埋まっていないレコードを探す。
	//なければリターン。
	//あればアップデートして、start_dateとの差を出して、リターン。

	
	res := domain.SlackResponse{Text: "end test", Channel: req.ChannelName, ResponseType: "in_channel"}
	return res
}
