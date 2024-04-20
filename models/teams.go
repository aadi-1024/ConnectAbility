package models

import "time"

type Teams struct {
	Id           int       `json:"id,omitempty" gorm:"primaryKey"`
	OwnerId      int       `json:"ownerId,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreationTime time.Time `json:"creationTime"`
}
