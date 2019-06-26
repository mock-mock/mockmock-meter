package utils

import (
	"github.com/labstack/echo"
	"github.com/mockmock-meter/backend/domain"
)

func ParseRequest(c echo.Context) domain.SlackRequest {
	return domain.SlackRequest{
		ChannelID: "test_id",
	}
}
