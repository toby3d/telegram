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

func (epe *EncryptedPassportElement) DecryptPassport(sv *SecureValue, b *Bot) (*Passport, error) {
	if !epe.IsPassport() || !sv.HasData() || !sv.HasFrontSide() {
		return nil, nil
	}

	/*
		var p Passport

		body, err := sv.Data.decrypt(epe.Data)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(body, &p.Data); err != nil {
			return nil, err
		}

		p.FrontSide, err = b.DecryptPassportFile(epe.FrontSide, sv.FrontSide)
		if err != nil {
			return nil, err
		}

		if sv.HasSelfie() {
			p.Selfie, err = b.DecryptPassportFile(epe.Selfie, sv.Selfie)
			if err != nil {
				return nil, err
			}
		}

		if sv.HasTranslation() {
			p.Translation = make([][]byte, len(sv.Translation))
			for i := range sv.Translation {
				p.Translation[i], err = b.DecryptPassportFile(epe.Translation[i], sv.Translation[i])
				if err != nil {
					return nil, err
				}
			}
		}
	*/

	return nil, nil
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
