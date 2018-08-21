package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetUserProfilePhotosParameters represents data for GetUserProfilePhotos method.
type GetUserProfilePhotosParameters struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// GetUserProfilePhotos get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (bot *Bot) GetUserProfilePhotos(params *GetUserProfilePhotosParameters) (photos *UserProfilePhotos, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetUserProfilePhotos)
	if err != nil {
		return
	}

	photos = new(UserProfilePhotos)
	err = json.Unmarshal(*resp.Result, photos)
	return
}
