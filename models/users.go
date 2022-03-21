package models

import (
	"golang.org/x/crypto/bcrypt"
)

const PasswordCryptLevel = 12

type Account struct {
	BaseModel
	Username string `gorm:"column:username;not null;unique_index;comment:'用户名'" json:"username" form:"username"`
	Password string `gorm:"column:password;comment:'密码'" form:"password" json:"-"`
	Name string `form:"name" json:"name"`
	IsActive bool `json:"-"`
	Phone string `gorm:"column:phone;not null;comment:'手机号'" json:"phone"`
	Email string `gorm:"column:email;not null;comment:'邮箱'" json:"email"`
}

func (a *Account) TableName() string {
	return "user_accounts"
}

func (a *Account) GetUserByID(id uint) *Account {
	DB.Model(&Account{}).First(a, id)
	if a.ID > 0 {
		return a
	} else {
		return nil
	}
}

// SetPassword 设置密码加密
func (a *Account) SetPassword(password string) error {
	p, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCryptLevel)
	if err != nil {
		return err
	}
	a.Password = string(p)
	return nil
}

// CheckPassword 验证登录帐户密码合法性
func (a *Account) CheckPassword() bool {
	password := a.Password
	DB.Where("username = ?", a.Username).First(&a)
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}

// IsPasswordEqual 验证登录帐户密码合法性
func (a *Account) IsPasswordEqual(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}

// CheckDuplicateUsername 验证用户名是否重复
func (a *Account) CheckDuplicateUsername() bool {
	var count int
	if DB.Model(&Account{}).Where("username=?", a.Username).Count(&count); count == 0 {
		// 用户名不存在
		return true
	} else {
		// 用户名已经存在
		return false
	}
}
