package models

import "time"

type TeamMember struct {
	TeamId      int       `json:"teamId,omitempty" gorm:"primaryKey,autoIncrement:false"`
	MemberId    int       `json:"memberId,omitempty" gorm:"primaryKey,autoIncrement:false"`
	JoiningDate time.Time `json:"joiningDate"`
	LeavingDate time.Time `json:"leavingDate"`
}
