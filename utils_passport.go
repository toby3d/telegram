package telegram

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
)

func decrypt(pk *rsa.PrivateKey, s, h, d string) (obj []byte, err error) {
	// Note that all base64-encoded fields should be decoded before use.
	secret, err := decodeField(s)
	if err != nil {
		return
	}

	hash, err := decodeField(h)
	if err != nil {
		return
	}

	data, err := decodeField(d)
	if err != nil {
		return
	}

	if pk != nil {
		// Decrypt the credentials secret (secret field in EncryptedCredentials)
		// using your private key
		secret, err = decryptSecret(pk, secret)
		if err != nil {
			return
		}
	}

	// Use this secret and the credentials hash (hash field in
	// EncryptedCredentials) to calculate credentials_key and credentials_iv
	key, iv := decryptSecretHash(secret, hash)
	if err != nil {
		return
	}

	// Decrypt the credentials data (data field in EncryptedCredentials) by
	// AES256-CBC using these credentials_key and credentials_iv.
	data, err = decryptData(key, iv, data)
	if err != nil {
		return
	}

	// IMPORTANT: At this step, make sure that the credentials hash is equal
	// to SHA256(credentials_data)
	if !match(hash, data) {
		err = ErrNotEqual
		return
	}

	// Credentials data is padded with 32 to 255 random padding bytes to make
	// its length divisible by 16 bytes. The first byte contains the length
	// of this padding (including this byte). Remove the padding to get the
	// data.
	offset := int(data[0])
	data = data[offset:]

	return
}

func decodeField(rawField string) (field []byte, err error) {
	return base64.StdEncoding.DecodeString(rawField)
}

func decryptSecret(pk *rsa.PrivateKey, s []byte) (secret []byte, err error) {
	return rsa.DecryptOAEP(sha1.New(), rand.Reader, pk, s, nil)
}

func decryptSecretHash(s, h []byte) (key, iv []byte) {
	hash := sha512.New()
	hash.Write(s)
	hash.Write(h)
	sh := hash.Sum(nil)

	return sh[0:32], sh[32 : 32+16]
}

func match(h, d []byte) bool {
	dh := sha256.New()
	dh.Write(d)

	return bytes.EqualFold(h, dh.Sum(nil))
}

func decryptData(key, iv, data []byte) (buf []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	buf = make([]byte, len(data))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(buf, data)

	return
}
