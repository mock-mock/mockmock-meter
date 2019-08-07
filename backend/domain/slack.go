package domain

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
