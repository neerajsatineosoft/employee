package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id               string `gorm: primarykey:"not null" json:"id"`
	Firstname        string `gorm:"size:50;not null" json:"first_name"`
	Lastname         string `gorm:"size:50;not null" json:"last_name"`
	Email            string `gorm:"size:50;not null" json:"email" binding:"`
	Password         string `gorm:"size:30;not null" json:"password"`
	Confirm_Password string `gorm:"size:30;not null" json:confirm_password`
	Phoneno          string `gorm:"size:20 not null" json:"phone_no"`
	Gender           string `gorm:"size:10 not null" json:"gender"`
}

func (u *User) Prepare() {
	u.Id = 0
	u.Firstname = html.EscapeString(strings.TrimSpace(u.Firstname))
	u.Lastname = html.EscapeString(strings.TrimSpace(u.Lastname))
	// 	u.From_date = time.Now()
	// 	u.To_date = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			return errors.New("Required email")
		}
		if u.Password == "" {
			return errors.New("Required password")
		}
	default:
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
	}
	return nil

}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	user := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&user).Error
	if err != nil {
		return &[]User{}, err
	}
	return &user, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"firstname": u.Firstname,
			"lastname":  u.Lastname,
			"email":     u.Email,
			"phone_no":  u.Phoneno,
			"gender":    u.Gender,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
