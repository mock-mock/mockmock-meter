package domain

type SlackRequest struct {

	// channel id
	ChannelID string `json:"channel_id,omitempty"`

	// channel name
	ChannelName string `json:"channel_name,omitempty"`

	// command
	Command string `json:"command,omitempty"`

	// response url
	ResponseURL string `json:"response_url,omitempty"`

	// team domain
	TeamDomain string `json:"team_domain,omitempty"`

	// team id
	TeamID string `json:"team_id,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// trigger id
	TriggerID string `json:"trigger_id,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`
}

type SlackResponse struct {

	// message
	Message string `json:"message,omitempty"`
}
