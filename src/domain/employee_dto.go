package domain

type Employee struct {
	id      uint64 `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key" json:"id"`
	name    string `json:"name"`
	email   string `json:"email"`
	address string `json:"address"`
}
