package telegram

import json "github.com/pquerna/ffjson/ffjson"

type SetPassportDataErrorsParameters struct {
	// User identifier
	UserID int `json:"user_id"`

	// A JSON-serialized array describing the errors
	Errors []PassportElementError `json:"errors"`
}

// SetPassportDataErrors informs a user that some of the Telegram Passport
// elements they provided contains errors. The user will not be able to re-submit
// their Passport to you until the errors are fixed (the contents of the field
// for which you returned the error must change). Returns True on success.
//
// Use this if the data submitted by the user doesn't satisfy the standards your
// service requires for any reason. For example, if a birthday date seems
// invalid, a submitted document is blurry, a scan shows evidence of tampering,
// etc. Supply some details in the error message to make sure the user knows how
// to correct the issues.
func (b *Bot) SetPassportDataErrors(userId int, errors []PassportElementError) (ok bool, err error) {
	dst, err := json.Marshal(&SetPassportDataErrorsParameters{
		UserID: userId,
		Errors: errors,
	})
	if err != nil {
		return
	}

	resp, err := b.request(dst, MethodSetPassportDataErrors)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
