package database

import (
	"github.com/aadi-1024/ConnectAbility/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *Database) RegisterUser(user *models.User) error {
	return d.conn.Create(user).Error
}

func (d *Database) LoginUser(user *models.User) (int, error) {
	userFetch := models.User{}
	res := d.conn.Table("users").Where("email = ?", user.Email).Select("id", "password").Scan(&userFetch)
	if res.Error != nil {
		return 0, nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(userFetch.Password), []byte(user.Password))
	if err != nil {
		return 0, err
	}
	return userFetch.Id, nil
}
