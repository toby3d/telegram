package telegram

import "time"

func (msg *Message) IsCommand() bool {
	if len(msg.Entities) <= 0 {
		return false
	}

	if msg.Entities[0].Type == EntityBotCommand &&
		msg.Entities[0].Offset == 0 {
		return true
	}

	return false
}

func (msg *Message) Command() string {
	if !msg.IsCommand() {
		return ""
	}

	return string([]rune(msg.Text)[1:msg.Entities[0].Length])
}

func (msg *Message) CommandArgument() string {
	if !msg.IsCommand() {
		return ""
	}

	if !msg.HasArgument() {
		return ""
	}

	return string([]rune(msg.Text)[(msg.Entities[0].Length + 1):])
}

func (msg *Message) HasArgument() bool {
	if !msg.IsCommand() {
		return false
	}

	if msg.CommandArgument() != "" {
		return true
	}

	return false
}

func (msg *Message) Time() time.Time {
	return time.Unix(msg.Date, 0)
}
