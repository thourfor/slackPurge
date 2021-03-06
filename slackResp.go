package main

import "encoding/json"

type slackResp struct {
	RespType    string   `json:"response_type,omitempty"`
	Text        string   `json:"text,omitempty"`
	Attachments []attach `json:"attachments,omitempty"`
}

type attach struct {
	Fallback    string   `json:"fallback,omitempty"`
	CallbackID  string   `json:"callback_id"`
	Color       string   `json:"color,omitempty"`
	Pretext     string   `json:"pretext,omitempty"`
	Author_name string   `json:"author_name,omitempty"`
	Author_link string   `json:"author_link,omitempty"`
	Author_icon string   `json:"author_icon,omitempty"`
	Title       string   `json:"title,omitempty"`
	Title_link  string   `json:"title_link,omitempty"`
	Text        string   `json:"text,omitempty"`
	Fields      []field  `json:"fields,omitempty"`
	Image_url   string   `json:"image_url,omitempty"`
	Thumb_url   string   `json:"thumb_url,omitempty"`
	Footer      string   `json:"footer,omitempty"`
	Footer_icon string   `json:"footer_icon,omitempty"`
	Ts          string   `json:"ts,omitempty"`
	Actions     []action `json:"actions,omitempty"`
}

type field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short string `json:"short,omitempty"`
}

type action struct {
	Name    string  `json:"name,omitempty"`
	Text    string  `json:"text,omitempty"`
	Style   string  `json:"style,omitempty"`
	Type    string  `json:"type,omitempty"`
	Value   string  `json:"value,omitempty"`
	Confirm confirm `json:"confirm,omitempty"`
}

type confirm struct {
	Title        string `json:"title,omitempty"`
	Text         string `json:"text,omitempty"`
	Ok_text      string `json:"ok_text,omitempty"`
	Dismiss_text string `json:"dismiss_text,omitempty"`
}

func (s *slackResp) String() string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return string(b)
}

type slackMsg struct {
	Payload slackPayload `json:"payload"`
}

type slackPayload struct {
	Actions      []action `json:"actions"`
	CallbackID   string   `json:"callback_id"`
	Team         teamInfo `json:"team"`
	Channel      channel  `json:"channel"`
	User         user     `json:"user"`
	ActionTS     string   `json:"action_ts"`
	MessageTS    string   `json:"message_ts"`
	AttachmentId string   `json:"attachment_id"`
	Token        string   `json:"token"`
	IsAppUnfurl  bool     `json:"is_app_unfurl"`
}

type teamInfo struct {
	ID     string `json"id"`
	Domain string `json:"domain"`
}

type channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
