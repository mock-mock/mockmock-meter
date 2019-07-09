package command

import (
	"fmt"
	"github.com/mock-mock/mockmock-meter/backend/domain"
)

// End Mock-Mock
func End(req domain.SlackRequest) domain.SlackResponse {
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

	res := domain.SlackResponse{Text: "end test"}
	return res
}
