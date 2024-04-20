package database

import (
	"errors"
	"github.com/aadi-1024/ConnectAbility/models"
	"gorm.io/gorm"
	"time"
)

func (d *Database) CreateTeam(teams models.Teams) (int, error) {
	err := d.conn.Transaction(func(tx *gorm.DB) error {
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
	res := d.conn.Table("teams").Where("Owner_Id = ?", uid).Limit(lim).Find(&ret)
	return ret, res.Error
}

func (d *Database) GetTeamById(uid, id int) (map[string]any, error) {
	team := models.Teams{}
	teamMember := models.TeamMember{}

	res := d.conn.Table("team_members").Where("Team_Id = ? and Member_Id = ?", id, uid).First(&teamMember)
	if res.RowsAffected == 0 {
		return nil, errors.New("unauthorized")
	}

	members := make([]int, 1)
	res = d.conn.Table("team_members").Where("Team_Id = ?", id).Select("Member_Id").Find(&members)

	if res.Error != nil {
		return nil, res.Error
	}

	err := d.conn.Find(&team, id).Error

	data := make(map[string]any)

	data["team"] = team
	data["members"] = members

	return data, err
}

func (d *Database) GetTeamOwner(id int) (int, error) {
	team := models.Teams{}
	res := d.conn.Select("Owner_Id").Find(&team, id)
	if res.Error != nil {
		return 0, res.Error
	}
	if res.RowsAffected == 0 {
		return 0, errors.New("invalid id")
	}
	return team.OwnerId, nil
}

func (d *Database) CreateTeamInvite(invite models.TeamInvite) error {
	return d.conn.Create(&invite).Error
}

func (d *Database) ResolveTeamInvite(uid, id int, accept bool) error {
	inv := models.TeamInvite{
		Id:        id,
		InvitedId: uid,
	}
	res := d.conn.Find(&inv)
	if res.Error != nil {
		return res.Error
	}
	if inv.TeamId == 0 {
		return errors.New("nuh uh")
	}
	tm := models.TeamMember{
		TeamId:      inv.TeamId,
		MemberId:    uid,
		JoiningDate: time.Now(),
	}
	err := d.conn.Transaction(func(tx *gorm.DB) error {
		if accept {
			err := tx.Omit("leaving_date").Create(&tm).Error
			if err != nil {
				return err
			}
		}
		if err := tx.Delete(&inv).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
