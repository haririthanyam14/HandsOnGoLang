package user

import (
	// "encoding/hex"
	// "encoding/json"
	"fmt"
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey,type:autoIncrement;" json:"-"`
	FirstName string    `type:"not null" json:"firstName"`
	LastName  string    `type:"not null" json:"lastName"`
	FullName  string    `type:"not null" json:"fullName"`
	UpdatedAt time.Time `gorm:"column:updatedat;default:now()" json:"-"`
	CreatedAt time.Time `gorm:"column:createdat;default:now()" json:"-"`
}

func NewUser(u UserRequest) (*User, error) {

	return &User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		FullName:  fmt.Sprintf("%s %s", u.FirstName, u.LastName),
	}, nil
}
