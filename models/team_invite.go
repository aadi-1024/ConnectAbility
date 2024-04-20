package models

type TeamInvite struct {
	Id        int `json:"id,omitempty" gorm:"primaryKey"`
	TeamId    int `json:"teamId,omitempty"`
	InvitedId int `json:"invitedIt,omitempty"`
}
