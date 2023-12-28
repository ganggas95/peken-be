package domain

type Role struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:""`
	Users       []User `gorm:"many2many:user_roles;" json:"users"`
}
