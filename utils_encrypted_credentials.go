package telegram

import (
	"crypto/rsa"

	json "github.com/pquerna/ffjson/ffjson"
)

func (ec *EncryptedCredentials) Decrypt(pk *rsa.PrivateKey) (*Credentials, error) {
	if ec == nil || pk == nil {
		return nil, nil
	}

	data, err := decrypt(pk, ec.Secret, ec.Hash, ec.Data)
	if err != nil {
		return nil, err
	}

	var c Credentials
	err = json.Unmarshal(data, &c)
	return &c, err
}
