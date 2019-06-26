package utils

import (
	"github.com/labstack/echo"

	slack "github.com/mock-mock/mockmock-meter/backend/domain"
)

func ParseRequest(c echo.Context) slack.SlackRequest {
	return slack.SlackRequest
}
