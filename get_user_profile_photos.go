package telegram

import json "github.com/pquerna/ffjson/ffjson"

type GetUserProfilePhotosParameters struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// GetUserProfilePhotos get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (bot *Bot) GetUserProfilePhotos(params *GetUserProfilePhotosParameters) (*UserProfilePhotos, error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodGetUserProfilePhotos)
	if err != nil {
		return nil, err
	}

	var data UserProfilePhotos
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
