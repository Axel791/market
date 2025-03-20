package domains

import "github.com/Axel791/appkit"

type LoyaltyBalance struct {
	ID     int64
	UserID int64
	Count  int64
}

func (v *LoyaltyBalance) ValidateUserID() error {
	if v.UserID <= 0 {
		return appkit.ValidationError("invalid userID")
	}
	return nil
}

func (v *LoyaltyBalance) ValidateCount() error {
	if v.Count < 0 {
		return appkit.ValidationError("invalid count")
	}
	return nil
}
