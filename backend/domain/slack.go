package domain

/**
Example of SlackRequest
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
type SlackRequest struct {

	// channel id
	ChannelID string `form:"channel_id"`

	// channel name
	ChannelName string `form:"channel_name"`

	// command
	Command string `form:"command"`

	// response url
	ResponseURL string `form:"response_url"`

	// team domain
	TeamDomain string `form:"team_domain"`

	// team id
	TeamID string `form:"team_id"`

	// text
	Text string `form:"text"`

	// token
	Token string `form:"token"`

	// trigger id
	TriggerID string `form:"trigger_id"`

	// user id
	UserID string `form:"user_id"`

	// user name
	UserName string `form:"user_name"`
}

// SlackResponse is ResponseDto
type SlackResponse struct {

	// message text
	Text string `json:"text,omitempty"`

	// return Channel
	Channel string `json:"channel,omitempty"`

	// https://api.slack.com/slash-commands
	ResponseType string `json:"response_type,omitempty"`
}
