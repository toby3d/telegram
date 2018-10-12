package telegram

import "errors"

var ErrNotEqual = errors.New("credentials hash and credentials data hash is not equal")

func (dc *DataCredentials) decrypt(d string) (data []byte, err error) {
	secret, err := decodeField(dc.Secret)
	if err != nil {
		return
	}

	hash, err := decodeField(dc.DataHash)
	if err != nil {
		return
	}

	key, iv := decryptSecretHash(secret, hash)
	if err != nil {
		return
	}

	data, err = decodeField(d)
	if err != nil {
		return
	}

	data, err = decryptData(key, iv, data)
	if err != nil {
		return
	}

	if !match(hash, data) {
		err = ErrNotEqual
	}

	offset := int(data[0])
	data = data[offset:]

	return
}
