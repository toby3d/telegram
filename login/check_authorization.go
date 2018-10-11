package login

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	http "github.com/valyala/fasthttp"
)

// CheckAuthorization verify the authentication and the integrity of the data
// received by comparing the received hash parameter with the hexadecimal
// representation of the HMAC-SHA-256 signature of the data-check-string with the
// SHA256 hash of the bot's token used as a secret key.
func CheckAuthorization(data interface{}, secretKey string) (bool, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)

	switch d := data.(type) {
	case *User:
		return d.CheckAuthorization(secretKey)
	case *http.Args:
		d.CopyTo(args)
		http.ReleaseArgs(d)
	case []byte:
		args.ParseBytes(d)
	case string:
		args.Parse(d)
	default:
		return false, ErrUnsupportedType

	}

	hash := args.Peek(KeyHash)
	args.Del(KeyHash)

	return check(args.QueryString(), []byte(secretKey), hash)
}

// CheckAuthorization verify the authentication and the integrity of the data
// received by comparing the received hash parameter with the hexadecimal
// representation of the HMAC-SHA-256 signature of the data-check-string with the
// SHA256 hash of the bot's token used as a secret key.
func (u *User) CheckAuthorization(secretKey string) (ok bool, err error) {
	args := u.toArgs()
	defer http.ReleaseArgs(args)
	hash := args.Peek(KeyHash)
	args.Del(KeyHash)

	return check(args.QueryString(), []byte(secretKey), hash)
}

func check(data, secretKey, hash []byte) (bool, error) {
	sk := sha256.Sum256(secretKey)
	h := hmac.New(sha256.New, sk[0:])
	if _, err := h.Write(data); err != nil {
		return false, err
	}

	return hex.EncodeToString(h.Sum(nil)) == string(hash), nil
}
