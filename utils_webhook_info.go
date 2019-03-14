package telegram

import "time"

func (wi *WebhookInfo) HasURL() bool {
	return wi != nil && wi.URL != ""
}

func (wi *WebhookInfo) LastErrorTime() *time.Time {
	if wi == nil {
		return nil
	}

	led := time.Unix(wi.LastErrorDate, 0)
	return &led
}
