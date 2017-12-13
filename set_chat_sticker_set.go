package telegram

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SetChatStickerSet set a new group sticker set for a supergroup. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Use the field
// can_set_sticker_set optionally returned in getChat requests to check if the bot can use this
// method. Returns True on success.
func (bot *Bot) SetChatStickerSet(chatID int64, stickerSetName string) (bool, error) {
	var args http.Args
	args.Add("chat_id", strconv.FormatInt(chatID, 10))
	args.Add("sticker_set_name", stickerSetName)

	resp, err := bot.request(nil, "setChatStickerSet", &args)
	if err != nil {
		return false, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
