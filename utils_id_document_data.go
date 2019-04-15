package telegram

import "time"

func (idd *IdDocumentData) ExpiryTime() *time.Time {
	if idd == nil || idd.ExpiryDate == "" {
		return nil
	}

	et, err := time.Parse("02.01.2006", idd.ExpiryDate)
	if err != nil {
		return nil
	}

	return &et
}
