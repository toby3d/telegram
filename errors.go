package telegram

import (
	"fmt"

	"golang.org/x/xerrors"
)

type Error struct {
	Code        int                   `json:"error_code"`
	Description string                `json:"description"`
	Parameters  []*ResponseParameters `json:"parameters,omitempty"`
	frame       xerrors.Frame
}

func (e Error) FormatError(p xerrors.Printer) error {
	p.Printf("%d %s", e.Code, e.Description)
	e.frame.Format(p)

	return nil
}

func (e Error) Format(s fmt.State, r rune) {
	xerrors.FormatError(e, s, r)
}

func (e Error) Error() string {
	return fmt.Sprint(e)
}
