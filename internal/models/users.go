package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Balance struct {
	Balance        float32 `json:"current"`
	TotalWithdrawn float32 `json:"withdrawn"`
}

type User struct {
	ID          int64     `json:"-"`
	Login       string    `json:"login,omitempty"`
	Password    []byte    `json:"-"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	Balan       Balance   `json:"-"`
}

func NewUser(login string, password string) *User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	return &User{
		Login:       login,
		Password:    hashedPassword,
		CreatedTime: time.Now(),
	}
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	return err == nil
}
