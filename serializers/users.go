package serializers

import "gin-fast/models"

type Register struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Phone string `form:"phone" json:"phone"`
	Email string `form:"email" json:"email"`
}

func (r *Register) GetUser() *models.Account {
	return &models.Account{
		Username: r.Username,
		Password: r.Password,
		Phone: r.Phone,
		Email: r.Email,
	}
}

type Login struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (l *Login) GetUser() *models.Account {
	return &models.Account{
		Username: l.Username,
		Password: l.Password,
	}
}

type Account struct {
	Username string `form:"username" json:"username"`
	OldPwd string `form:"oldPwd" json:"oldPwd"`
	NewPwd string `form:"newPwd" json:"newPwd"`
}
