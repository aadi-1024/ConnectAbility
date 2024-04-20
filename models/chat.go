package models

type Chat struct {
	Id      int `json:"id,omitempty" gorm:"primaryKey,autoIncrement:false"`
	Member1 int `json:"member1,omitempty"`
	Member2 int `json:"member2,omitempty"`
}
