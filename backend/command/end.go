package command

import "github.com/mock-mock/mockmock-meter/backend/domain"

func End(req domain.SlackRequest) string {
	return "end"
}
