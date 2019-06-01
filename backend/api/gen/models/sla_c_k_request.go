// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SLACKRequest slack request
// swagger:model SlackRequest
type SLACKRequest struct {

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

// Validate validates this slack request
func (m *SLACKRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SLACKRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SLACKRequest) UnmarshalBinary(b []byte) error {
	var res SLACKRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}