package database

import "github.com/aadi-1024/ConnectAbility/models"

func (d *Database) RegisterUser(user *models.User) error {
	return d.Conn.Create(user).Error
}
