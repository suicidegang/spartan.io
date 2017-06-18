package main

import (
	"strings"
	"encoding/json"
	"time"
)

type PackedMessage struct {
	Channel string
	Message map[string]interface{}
}

type Event struct {
	Room    string                 `json:"room"`
	Event   string                 `json:"event"`
	Message map[string]interface{} `json:"message"`
}

func (event Event) RoomID() string {
	return strings.Replace(event.Event, " ", ":", -1)
}

func (event Event) Encode() string {
	bytes, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}


// Pack a list of messages.
func list(messages ...map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"list": messages,
	}
}

// Anonymous message builder.
func anonymousMessage(str string) map[string]interface{} {
	str = strings.TrimSpace(str)

	if len(str) >= 200 {
		str = str[:200] + "..."
	}

	return map[string]interface{}{
		"content":        str,
		"user_id":        "guest",
		"username":       "guest",
		"avatar":         false,
		"timestamp":      time.Now().Unix(),
		"timestamp_nano": time.Now().UnixNano(),
	}
}

func userMessage(user *User, str string) map[string]interface{} {
	str = strings.TrimSpace(str)

	if len(str) >= 200 {
		str = str[:200] + "..."
	}

	return map[string]interface{}{
		"content":        str,
		"user_id":        user.Id,
		"username":       user.UserName,
		"avatar":         user.Image,
		"timestamp":      time.Now().Unix(),
		"timestamp_nano": time.Now().UnixNano(),
	}
}