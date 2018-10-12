package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetUserProfilePhotosParameters represents data for GetUserProfilePhotos method.
type GetUserProfilePhotosParameters struct {
	// Unique identifier of the target user
	UserID int `json:"user_id"`

	// Sequential number of the first photo to be returned. By default, all
	// photos are returned.
	Offset int `json:"offset,omitempty"`

	// Limits the number of photos to be retrieved. Values between 1—100 are
	// accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
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
