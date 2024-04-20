package database

import (
	"errors"
	"github.com/aadi-1024/ConnectAbility/models"
	"gorm.io/gorm"
	"time"
)

func (d *Database) CreateTeam(teams models.Teams) (int, error) {
	err := d.Conn.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&teams).Error
		if err != nil {
			return err
		}

		tm := models.TeamMember{
			TeamId:      teams.Id,
			MemberId:    teams.OwnerId,
			JoiningDate: time.Now(),
		}

		err = tx.Table("team_members").Omit("leaving_date").Create(&tm).Error
		if err != nil {
			return err
		}
		return err
	})
	return teams.Id, err
}

func (d *Database) GetTeams(uid, lim int) ([]*models.Teams, error) {
	ret := make([]*models.Teams, lim)
	res := d.Conn.Table("teams").Where("Owner_Id = ?", uid).Limit(lim).Find(&ret)
	return ret, res.Error
}

func (d *Database) GetTeamById(uid, id int) (models.Teams, error) {
	team := models.Teams{}
	teamMember := models.TeamMember{}
	res := d.Conn.Table("team_members").Where("Team_Id = ? and Member_Id = ?", id, uid).First(&teamMember)
	if res.RowsAffected == 0 {
		return team, errors.New("unauthorized")
	}
	err := d.Conn.Find(&team, id).Error
	return team, err
}
