package models

type User struct {
	Id              int    `json:"id,omitempty" gorm:"primaryKey,autoIncrement:false"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	PhoneNo         string `json:"phoneNo,omitempty"`
	FirstName       string `json:"firstName,omitempty"`
	LastName        string `json:"lastName,omitempty"`
	About           string `json:"about,omitempty"`
	ProfilePic      string `json:"profilePic,omitempty"`
	ResumeLink      string `json:"resumeLink,omitempty"`
	Github          string `json:"github,omitempty"`
	Linkedin        string `json:"linkedin,omitempty"`
	Website         string `json:"website,omitempty"`
	LocationArea    string `json:"locationArea,omitempty"`
	LocationCity    string `json:"locationCity,omitempty"`
	LocationCountry string `json:"locationCountry,omitempty"`
	LocationPin     string `json:"locationPin,omitempty"`
}
