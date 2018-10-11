package login

import http "github.com/valyala/fasthttp"

// ParseUser create User structure from input url.Values.
func ParseUser(data interface{}) (*User, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)

	switch d := data.(type) {
	case *http.Args:
		d.CopyTo(args)
		http.ReleaseArgs(d)
	case []byte:
		args.ParseBytes(d)
	case string:
		args.Parse(d)
	default:
		return nil, ErrUnsupportedType
	}

	return &User{
		ID:        args.GetUintOrZero(KeyID),
		AuthDate:  int64(args.GetUintOrZero(KeyAuthDate)),
		FirstName: string(args.Peek(KeyFirstName)),
		Hash:      string(args.Peek(KeyHash)),
		LastName:  string(args.Peek(KeyLastName)),
		PhotoURL:  string(args.Peek(KeyPhotoURL)),
		Username:  string(args.Peek(KeyUsername)),
	}, nil
}
