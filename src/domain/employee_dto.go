package domain

type Employee struct {
	id      int64  `gorm:"auto_increment;not null;primary_key" json:"id"`
	name    string `json:"name"`
	email   string `json:"email"`
	address string `json:"address"`
}
