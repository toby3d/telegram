package telegram

import (
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

// parseResponseError unmarshal src bytes into dst or return response Error.
func parseResponseError(marshler jsoniter.API, src []byte, dst interface{}) (err error) {
	resp := new(Response)
	if err = marshler.Unmarshal(src, resp); err != nil {
		return
	}

	if resp.Ok {
		return marshler.Unmarshal(resp.Result, dst)
	}

	respErr := new(Error)
	respErr.Code = resp.ErrorCode
	respErr.Description = resp.Description
	respErr.frame = xerrors.Caller(1)

	copy(respErr.Parameters, resp.Parameters)

	return respErr
}
