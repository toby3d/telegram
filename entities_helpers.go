package telegram

import "net/url"

func (entity *MessageEntity) ParseURL() (*url.URL, error) {
	if entity.Type == EntityTextLink {
		return url.Parse(entity.URL)
	}

	return nil, nil
}
