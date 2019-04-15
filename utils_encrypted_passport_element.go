package telegram

import (
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
)

func (epe *EncryptedPassportElement) DecryptPersonalDetails(sv *SecureValue) (*PersonalDetails, error) {
	if !epe.IsPersonalDetails() || !sv.HasData() {
		return nil, nil
	}

	body, err := sv.Data.decrypt(epe.Data)
	if err != nil {
		return nil, err
	}

	var pd PersonalDetails
	err = json.Unmarshal(body, &pd)
	return &pd, err
}

func (epe *EncryptedPassportElement) DecryptPassport(sv *SecureValue, b *Bot) (*IdDocumentData, []byte, []byte, [][]byte, error) {
	if !epe.IsPassport() || !sv.HasData() || !sv.HasFrontSide() {
		return nil, nil, nil, nil, nil
	}

	body, err := sv.Data.decrypt(epe.Data)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var idd IdDocumentData
	if err = json.Unmarshal(body, &idd); err != nil {
		return nil, nil, nil, nil, err
	}

	fs, err := b.DecryptFile(epe.FrontSide, sv.FrontSide)
	if err != nil {
		return &idd, nil, nil, nil, err
	}

	var s []byte
	if sv.HasSelfie() {
		if s, err = b.DecryptFile(epe.Selfie, sv.Selfie); err != nil {
			return &idd, fs, nil, nil, err
		}
	}

	t := make([][]byte, len(sv.Translation))
	if sv.HasTranslation() {
		for i := range t {
			if t[i], err = b.DecryptFile(&epe.Translation[i], &sv.Translation[i]); err != nil {
				return &idd, fs, s, nil, err
			}
		}
	}

	return &idd, fs, s, t, nil
}

func (epe *EncryptedPassportElement) DecryptInternalPassport(sv *SecureValue, b *Bot) (*IdDocumentData, []byte, []byte, [][]byte, error) {
	if !epe.IsInternalPassport() || !sv.HasData() || !sv.HasFrontSide() {
		return nil, nil, nil, nil, nil
	}

	body, err := sv.Data.decrypt(epe.Data)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var idd IdDocumentData
	if err = json.Unmarshal(body, &idd); err != nil {
		return nil, nil, nil, nil, err
	}

	fs, err := b.DecryptFile(epe.FrontSide, sv.FrontSide)
	if err != nil {
		return &idd, nil, nil, nil, err
	}

	var s []byte
	if sv.HasSelfie() {
		if s, err = b.DecryptFile(epe.Selfie, sv.Selfie); err != nil {
			return &idd, fs, nil, nil, err
		}
	}

	t := make([][]byte, len(sv.Translation))
	if sv.HasTranslation() {
		for i := range t {
			if t[i], err = b.DecryptFile(&epe.Translation[i], &sv.Translation[i]); err != nil {
				return &idd, fs, s, nil, err
			}
		}
	}

	return &idd, fs, s, t, nil
}

func (epe *EncryptedPassportElement) DecryptDriverLicense(sv *SecureValue, b *Bot) (*IdDocumentData, []byte, []byte, []byte, [][]byte, error) {
	if !epe.IsDriverLicense() || !sv.HasData() || !sv.HasFrontSide() || !sv.HasReverseSide() {
		return nil, nil, nil, nil, nil, nil
	}

	body, err := sv.Data.decrypt(epe.Data)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	var idd IdDocumentData
	if err = json.Unmarshal(body, &idd); err != nil {
		return nil, nil, nil, nil, nil, err
	}

	fs, err := b.DecryptFile(epe.FrontSide, sv.FrontSide)
	if err != nil {
		return &idd, nil, nil, nil, nil, err
	}

	rs, err := b.DecryptFile(epe.ReverseSide, sv.ReverseSide)
	if err != nil {
		return &idd, nil, nil, nil, nil, err
	}

	var s []byte
	if sv.HasSelfie() {
		if s, err = b.DecryptFile(epe.Selfie, sv.Selfie); err != nil {
			return &idd, fs, rs, nil, nil, err
		}
	}

	t := make([][]byte, len(sv.Translation))
	if sv.HasTranslation() {
		for i := range t {
			if t[i], err = b.DecryptFile(&epe.Translation[i], &sv.Translation[i]); err != nil {
				return &idd, fs, rs, s, nil, err
			}
		}
	}

	return &idd, fs, rs, s, t, nil
}

func (epe *EncryptedPassportElement) IsAddress() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeAddress)
}

func (epe *EncryptedPassportElement) IsBankStatement() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeBankStatement)
}

func (epe *EncryptedPassportElement) IsDriverLicense() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeDriverLicense)
}

func (epe *EncryptedPassportElement) IsEmail() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeEmail)
}

func (epe *EncryptedPassportElement) IsIdentityCard() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeIdentityCard)
}

func (epe *EncryptedPassportElement) IsInternalPassport() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeInternalPassport)
}

func (epe *EncryptedPassportElement) IsPassport() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypePassport)
}

func (epe *EncryptedPassportElement) IsPassportRegistration() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypePassportRegistration)
}

func (epe *EncryptedPassportElement) IsPersonalDetails() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypePersonalDetails)
}

func (epe *EncryptedPassportElement) IsPhoneNumber() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypePhoneNumber)
}

func (epe *EncryptedPassportElement) IsRentalAgreement() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeRentalAgreement)
}

func (epe *EncryptedPassportElement) IsTemporaryRegistration() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeTemporaryRegistration)
}

func (epe *EncryptedPassportElement) IsUtilityBill() bool {
	return epe != nil && strings.EqualFold(epe.Type, TypeUtilityBill)
}
