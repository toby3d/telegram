package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetUserProfilePhotos get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (bot *Bot) GetUserProfilePhotos(userID, offset, limit int) (*UserProfilePhotos, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("user_id", strconv.Itoa(userID))
	args.Add("offset", strconv.Itoa(offset))

	if limit > 0 {
		args.Add("limit", strconv.Itoa(limit))
	}

	resp, err := bot.request(nil, "getUserProfilePhotos", args)
	if err != nil {
		return nil, err
	}

	var data UserProfilePhotos
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}
