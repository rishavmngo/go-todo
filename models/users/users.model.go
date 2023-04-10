package userModel

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rishavmngo/todo/db"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Age      int32  `json:"age"`
	Email    string `gorm:"unique;not null;default:null" json:"email"`
	Password string `gorm:"not null;default:null" json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (user *User) Register() (*User, error) {
	db.NewRecord(user)

	if result := db.Create(user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (user *User) GetByEmail() (*User, error) {
	if result := db.Where("email = ?", user.Email).Find(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
func (user *User) Validate(action string) error {

	switch action {
	case "login":
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}

	}
	return nil

}
