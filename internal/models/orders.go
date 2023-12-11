package models

import "time"

type Order struct {
	ID         int64     `json:"-"`
	UserID     int64     `json:"-"`
	Number     string    `json:"number,omitempty"`
	Status     string    `json:"status,omitempty"`
	Accrual    float32   `json:"accrual,omitempty"`
	UploadedAt time.Time `json:"uploaded_at,omitempty"`
}

func NewOrder(userID int64, number string, status string, accrual float32) *Order {
	return &Order{
		UserID:     userID,
		Number:     number,
		Status:     status,
		Accrual:    accrual,
		UploadedAt: time.Now(),
	}
}
