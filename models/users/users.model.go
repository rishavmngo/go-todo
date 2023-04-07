package userModel

import (
	"github.com/jinzhu/gorm"
	"github.com/rishavmngo/todo/db"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int32  `json:"age"`
	Email string `gorm:"unique;not null;default:null" json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (user *User) Register() (*User, error) {
	// db.NewRecord(user)

	if result := db.Create(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
