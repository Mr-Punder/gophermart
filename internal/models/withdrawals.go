package models

import "time"

type Withdrawal struct {
	ID          int64     `json:"-"`
	UserID      int64     `json:"-"`
	OrderNumber int64     `json:"order,omitempty"`
	Amount      float32   `json:"sum,omitempty"`
	ProcessedAt time.Time `json:"processed_at,omitempty"`
}

func NewWithdrawal(userID int64, number int64, amount float32) *Withdrawal {
	return &Withdrawal{
		UserID:      userID,
		OrderNumber: number,
		Amount:      amount,
	}
}
