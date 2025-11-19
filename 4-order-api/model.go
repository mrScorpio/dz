package orderapi

import (
	"github.com/lib/pq"
	concurrency "github.com/mrScorpio/dz/1-concurrency"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description"`
	PriceRub    int            `json:"prub" validate:"required"`
	Images      pq.StringArray `json:"images" gorm:"type:varchar(256)[]"`
}

func NewProduct(name, descr string, prub int) *Product {
	return &Product{
		Name:        name,
		Description: descr,
		PriceRub:    prub,
	}
}

type User struct {
	gorm.Model
	Phone     string `json:"phone" validate:"required"`
	SessionID string `json:"session_id"`
	Secret    string `json:"secret"`
	Code      string `json:"code"`
}

func NewUser(phone string) *User {
	user := &User{
		Phone: phone,
	}
	user.GenSessionId()
	user.GenSecret()
	return user
}

func (u *User) GenSessionId() {
	idNum := concurrency.GenNums(16, 26)
	idR := make([]rune, len(idNum))
	for i, v := range idNum {
		idR[i] = rune(97 + v)
	}
	u.SessionID = string(idR)
	u.GenCode()
}

func (u *User) GenSecret() {
	idNum := concurrency.GenNums(32, 26)
	idR := make([]rune, len(idNum))
	for i, v := range idNum {
		idR[i] = rune(97 + v)
	}
	u.Secret = string(idR)
}

func (u *User) GenCode() {
	idNum := concurrency.GenNums(6, 9)
	idR := make([]rune, len(idNum))
	for i, v := range idNum {
		idR[i] = rune(49 + v)
	}
	u.Code = string(idR)
}
