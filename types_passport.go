package telegram

type (
	Fields struct {
		PersonalDetails       *PersonalDetails
		Passport              *Passport
		InternalPassport      *InternalPassport
		DriverLicense         *DriverLicense
		IdentityCard          *IdentityCard
		Address               *ResidentialAddress
		UtilityBill           *UtilityBill
		BankStatement         *BankStatement
		RentalAgreement       *RentalAgreement
		PassportRegistration  *PassportRegistration
		TemporaryRegistration *TemporaryRegistration
		PhoneNumber           PhoneNumber
		Email                 Email
	}

	// Passport represent passport.
	Passport struct {
		Data        *IdDocumentData
		FrontSide   *PassportFile
		Selfie      *PassportFile
		Translation []PassportFile
	}

	// InternalPassport represent internal passport.
	InternalPassport struct {
		Data        *IdDocumentData
		FrontSide   *PassportFile
		Selfie      *PassportFile
		Translation []PassportFile
	}

	// DriverLicense represent driver license.
	DriverLicense struct {
		Passport
		ReverseSide *PassportFile
	}

	// IdentityCard represent identity card.
	IdentityCard struct {
		Data        *IdDocumentData
		FrontSide   *PassportFile
		ReverseSide *PassportFile
		Selfie      *PassportFile
		Translation []PassportFile
	}

	// UtilityBill represent utility bill.
	UtilityBill struct {
		Files       []PassportFile
		Translation []PassportFile
	}

	// BankStatement represent bank statement.
	BankStatement struct {
		Files       []PassportFile
		Translation []PassportFile
	}

	// RentalAgreement represent rental agreement.
	RentalAgreement struct {
		Files       []PassportFile
		Translation []PassportFile
	}

	// PassportRegistration represent registration Page in the internal passport.
	PassportRegistration struct {
		Files       []PassportFile
		Translation []PassportFile
	}

	// TemporaryRegistration represent temporary registration.
	TemporaryRegistration struct {
		Files       []PassportFile
		Translation []PassportFile
	}

	// PhoneNumber represent phone number.
	PhoneNumber string

	// Email represent email.
	Email string

	// PersonalDetails represents personal details.
	PersonalDetails struct {
		// First Name
		FirstName string `json:"first_name"`

		// Last Name
		LastName string `json:"last_name"`

		// Middle Name
		MiddleName string `json:"middle_name,omitempty"`

		// Date of birth in DD.MM.YYYY format
		BirthDate string `json:"birth_date"`

		// Gender, male or female
		Gender string `json:"gender"`

		// Citizenship (ISO 3166-1 alpha-2 country code)
		CountryCode string `json:"country_code"`

		// Country of residence (ISO 3166-1 alpha-2 country code)
		ResidenceCountryCode string `json:"residence_country_code"`

		// First Name in the language of the user's country of residence
		FirstNameNative string `json:"first_name_native"`

		// Last Name in the language of the user's country of residence
		LastNameNative string `json:"last_name_native"`

		// Middle Name in the language of the user's country of residence
		MiddleNameNative string `json:"middle_name_native,omitempty"`
	}

	// ResidentialAddress represents a residential address.
	ResidentialAddress struct {
		// First line for the address
		StreetLine1 string `json:"street_line1"`

		// Second line for the address
		StreetLine2 string `json:"street_line2,omitempty"`

		// City
		City string `json:"city"`

		// State
		State string `json:"state,omitempty"`

		// ISO 3166-1 alpha-2 country code
		CountryCode string `json:"country_code"`

		// Address post code
		PostCode string `json:"post_code"`
	}

	// IdDocumentData represents the data of an identity document.
	IdDocumentData struct {
		// Document number
		DocumentNo string `json:"document_no"`

		// Date of expiry, in DD.MM.YYYY format
		ExpiryDate string `json:"expiry_date,omitempty"`
	}

	// PassportData contains information about Telegram Passport data shared with
	// the bot by the user.
	PassportData struct {
		// Array with information about documents and other Telegram Passport
		// elements that was shared with the bot
		Data []EncryptedPassportElement `json:"data"`

		// Encrypted credentials required to decrypt the data
		Credentials *EncryptedCredentials `json:"credentials"`
	}

	// PassportFile represents a file uploaded to Telegram Passport. Currently all
	// Telegram Passport files are in JPEG format when decrypted and don't exceed
	// 10MB.
	PassportFile struct {
		// Unique identifier for this file
		FileID string `json:"file_id"`

		// File size
		FileSize int `json:"file_size"`

		// Unix time when the file was uploaded
		FileDate int64 `json:"file_date"`
	}

	// Credentials is a JSON-serialized object.
	Credentials struct {
		// Credentials for encrypted data
		SecureData *SecureData `json:"secure_data"`

		// Bot-specified nonce
		Nonce string `json:"nonce"`
	}

	// SecureData represents the credentials required to decrypt encrypted
	// data. All fields are optional and depend on fields that were requested.
	SecureData struct {
		// Credentials for encrypted personal details
		PersonalDetails *SecureValue `json:"personal_details,omitempty"`

		// Credentials for encrypted passport
		Passport *SecureValue `json:"passport,omitempty"`

		// Credentials for encrypted internal passport
		InternalPassport *SecureValue `json:"internal_passport,omitempty"`

		// Credentials for encrypted driver license
		DriverLicense *SecureValue `json:"driver_license,omitempty"`

		// Credentials for encrypted ID card
		IdentityCard *SecureValue `json:"identity_card,omitempty"`

		// Credentials for encrypted residential address
		Address *SecureValue `json:"address,omitempty"`

		// Credentials for encrypted utility bill
		UtilityBill *SecureValue `json:"utility_bill,omitempty"`

		// Credentials for encrypted bank statement
		BankStatement *SecureValue `json:"bank_statement,omitempty"`

		// Credentials for encrypted rental agreement
		RentalAgreement *SecureValue `json:"rental_agreement,omitempty"`

		// Credentials for encrypted registration from internal passport
		PassportRegistration *SecureValue `json:"passport_registration,omitempty"`

		// Credentials for encrypted temporary registration
		TemporaryRegistration *SecureValue `json:"temporary_registration,omitempty"`
	}

	// SecureValue represents the credentials required to decrypt encrypted
	// values. All fields are optional and depend on the type of fields that
	// were requested.
	SecureValue struct {
		// Credentials for encrypted Telegram Passport data.
		Data *DataCredentials `json:"data,omitempty"`

		// Credentials for an encrypted document's front side.
		FrontSide *FileCredentials `json:"front_side,omitempty"`

		// Credentials for an encrypted document's reverse side.
		ReverseSide *FileCredentials `json:"reverse_side,omitempty"`

		// Credentials for an encrypted selfie of the user with a document.
		Selfie *FileCredentials `json:"selfie,omitempty"`

		// Credentials for an encrypted translation of the document.
		Translation []FileCredentials `json:"translation,omitempty"`

		// Credentials for encrypted files.
		Files []FileCredentials `json:"files,omitempty"`
	}

	// DataCredentials can be used to decrypt encrypted data from the data
	// field in EncryptedPassportElement.
	DataCredentials struct {
		// Checksum of encrypted data
		DataHash string `json:"data_hash"`

		// Secret of encrypted data
		Secret string `json:"secret"`
	}

	// FileCredentials can be used to decrypt encrypted files from the
	// front_side, reverse_side, selfie, files and translation fields in
	// EncryptedPassportElement.
	FileCredentials struct {
		// Checksum of encrypted file
		FileHash string `json:"file_hash"`

		// Secret of encrypted file
		Secret string `json:"secret"`
	}

	// EncryptedPassportElement contains information about documents or other
	// Telegram Passport elements shared with the bot by the user.
	EncryptedPassportElement struct {
		// Element type.
		Type string `json:"type"`

		// Base64-encoded encrypted Telegram Passport element data provided by
		// the user, available for "personal_details", "passport",
		// "driver_license", "identity_card", "identity_passport" and "address"
		// types. Can be decrypted and verified using the accompanying
		// EncryptedCredentials.
		Data string `json:"data,omitempty"`

		// User's verified phone number, available only for "phone_number" type
		PhoneNumber string `json:"phone_number,omitempty"`

		// User's verified email address, available only for "email" type
		Email string `json:"email,omitempty"`

		// Array of encrypted files with documents provided by the user,
		// available for "utility_bill", "bank_statement", "rental_agreement",
		// "passport_registration" and "temporary_registration" types. Files can
		// be decrypted and verified using the accompanying EncryptedCredentials.
		Files []PassportFile `json:"files,omitempty"`

		// Encrypted file with the front side of the document, provided by the
		// user. Available for "passport", "driver_license", "identity_card" and
		// "internal_passport". The file can be decrypted and verified using the
		// accompanying EncryptedCredentials.
		FrontSide *PassportFile `json:"front_side,omitempty"`

		// Encrypted file with the reverse side of the document, provided by the
		// user. Available for "driver_license" and "identity_card". The file can
		// be decrypted and verified using the accompanying EncryptedCredentials.
		ReverseSide *PassportFile `json:"reverse_side,omitempty"`

		// Encrypted file with the selfie of the user holding a document,
		// provided by the user; available for "passport", "driver_license",
		// "identity_card" and "internal_passport". The file can be decrypted
		// and verified using the accompanying EncryptedCredentials.
		Selfie *PassportFile `json:"selfie,omitempty"`

		// Array of encrypted files with translated versions of documents
		// provided by the user. Available if requested for “passport”,
		// “driver_license”, “identity_card”, “internal_passport”,
		// “utility_bill”, “bank_statement”, “rental_agreement”,
		// “passport_registration” and “temporary_registration” types.
		// Files can be decrypted and verified using the accompanying
		// EncryptedCredentials.
		Translation []PassportFile `json:"translation,omitempty"`

		// Base64-encoded element hash for using in PassportElementErrorUnspecified
		Hash string `json:"hash"`
	}

	// EncryptedCredentials contains data required for decrypting and
	// authenticating EncryptedPassportElement. See the Telegram Passport
	// Documentation for a complete description of the data decryption and
	// authentication processes.
	EncryptedCredentials struct {
		// Base64-encoded encrypted JSON-serialized data with unique user's
		// payload, data hashes and secrets required for EncryptedPassportElement
		// decryption and authentication
		Data string `json:"data"`

		// Base64-encoded data hash for data authentication
		Hash string `json:"hash"`

		// Base64-encoded secret, encrypted with the bot's public RSA key,
		// required for data decryption
		Secret string `json:"secret"`
	}

	// PassportElementError represents an error in the Telegram Passport element
	// which was submitted that should be resolved by the user.
	PassportElementError interface {
		PassportElementErrorMessage() string
		PassportElementErrorSource() string
		PassportElementErrorType() string
	}

	// PassportElementErrorDataField represents an issue in one of the data
	// fields that was provided by the user. The error is considered resolved
	// when the field's value changes.
	PassportElementErrorDataField struct {
		// Error source, must be data
		Source string `json:"source"`

		// The section of the user's Telegram Passport which has the error, one
		// of "personal_details", "passport", "driver_license", "identity_card",
		// "internal_passport", "address"
		Type string `json:"type"`

		// Name of the data field which has the error
		FieldName string `json:"field_name"`

		// Base64-encoded data hash
		DataHash string `json:"data_hash"`

		// Error message
		Message string `json:"message"`
	}

	// PassportElementErrorFrontSide represents an issue with the front side of
	// a document. The error is considered resolved when the file with the front
	// side of the document changes.
	PassportElementErrorFrontSide struct {
		// Error source, must be front_side
		Source string `json:"source"`

		// The section of the user's Telegram Passport which has the issue, one
		// of "passport", "driver_license", "identity_card", "internal_passport"
		Type string `json:"type"`

		// Base64-encoded hash of the file with the front side of the document
		FileHash string `json:"file_hash"`

		// Error message
		Message string `json:"message"`
	}

	// PassportElementErrorReverseSide represents an issue with the reverse side
	// of a document. The error is considered resolved when the file with reverse
	// side of the document changes.
	PassportElementErrorReverseSide struct {
		// Error source, must be reverse_side
		Source string `json:"source"`

		// The section of the user's Telegram Passport which has the issue, one
		// of "driver_license", "identity_card"
		Type string `json:"type"`

		// Base64-encoded hash of the file with the reverse side of the document
		FileHash string `json:"file_hash"`

		// Error message
		Message string `json:"message"`
	}

	// PassportElementErrorSelfie represents an issue with the selfie with a
	// document. The error is considered resolved when the file with the selfie
	// changes.
	PassportElementErrorSelfie struct {
		// Error source, must be selfie
		Source string `json:"source"`

		// The section of the user's Telegram Passport which has the issue, one
		// of "passport", "driver_license", "identity_card", "internal_passport"
		Type string `json:"type"`

		// Base64-encoded hash of the file with the selfie
		FileHash string `json:"file_hash"`

		// Error message
		Message string `json:"message"`
	}

	// PassportElementErrorFile represents an issue with a document scan. The
	// error is considered resolved when the file with the document scan changes.
	PassportElementErrorFile struct {
		// Error source, must be file
		Source string `json:"source"`

		// The section of the user's Telegram Passport which has the issue, one
		// of "utility_bill", "bank_statement", "rental_agreement",
		// "passport_registration", "temporary_registration"
		Type string `json:"type"`

		// Base64-encoded file hash
		FileHash string `json:"file_hash"`

		// Error message
		Message string `json:"message"`
	}

	// PassportElementErrorFiles represents an issue with a list of scans. The
	// error is considered resolved when the list of files containing the scans
	// changes.
	PassportElementErrorFiles struct {
		// Error source, must be files
		Source string `json:"source"`

		// The section of the user's Telegram Passport which has the issue, one
		// of "utility_bill", "bank_statement", "rental_agreement",
		// "passport_registration", "temporary_registration"
		Type string `json:"type"`

		// List of base64-encoded file hashes
		FileHashes []string `json:"file_hashes"`

		// Error message
		Message string `json:"message"`
	}

	// PassportElementErrorTranslationFile represents an issue with one of the
	// files that constitute the translation of a document. The error is
	// considered resolved when the file changes.
	PassportElementErrorTranslationFile struct {
		// Error source, must be translation_file
		Source string `json:"source"`

		// Type of element of the user's Telegram Passport which has the issue,
		// one of “passport”, “driver_license”, “identity_card”,
		// “internal_passport”, “utility_bill”, “bank_statement”,
		// “rental_agreement”, “passport_registration”, “temporary_registration”
		Type string `json:"type"`

		// Base64-encoded file hash
		FileHash string `json:"file_hash"`

		// Error message
		Message string `json:"message"`
	}

	// PassportElementErrorTranslationFiles represents an issue with the translated
	// version of a document. The error is considered resolved when a file with the
	// document translation change.
	PassportElementErrorTranslationFiles struct {
		// Error source, must be translation_files
		Source string `json:"source"`

		// Type of element of the user's Telegram Passport which has the issue,
		// one of “passport”, “driver_license”, “identity_card”,
		// “internal_passport”, “utility_bill”, “bank_statement”,
		// “rental_agreement”, “passport_registration”, “temporary_registration”
		Type string `json:"type"`

		// List of base64-encoded file hashes
		FileHashes []string `json:"file_hashes"`

		// Error message
		Message string `json:"message"`
	}

	// PassportElementErrorUnspecified represents an issue in an unspecified place.
	// The error is considered resolved when new data is added.
	PassportElementErrorUnspecified struct {
		// Error source, must be unspecified
		Source string `json:"source"`

		// Type of element of the user's Telegram Passport which has the issue
		Type string `json:"type"`

		// Base64-encoded element hash
		ElementHash string `json:"element_hash"`

		// Error message
		Message string `json:"message"`
	}
)
