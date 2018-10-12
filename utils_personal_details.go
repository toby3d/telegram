package telegram

import "time"

func (pd *PersonalDetails) BirthTime() *time.Time {
	if pd == nil || pd.BirthDate == "" {
		return nil
	}

	bt, err := time.Parse("02.01.2006", pd.BirthDate)
	if err != nil {
		return nil
	}

	return &bt
}

func (pd *PersonalDetails) FullName() string {
	if pd == nil {
		return ""
	}

	return pd.FirstName + " " + pd.LastName
}

func (pd *PersonalDetails) FullNameNative() string {
	if pd == nil {
		return ""
	}

	return pd.FirstNameNative + " " + pd.LastNameNative
}
