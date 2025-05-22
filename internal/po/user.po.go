package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"colum:uuid; type:varchar(255); not null; index:idx_uuid; unique;"`
	UserName string    `gorm:"colum:user_name"`
	IsActive bool      `gorm:"colum:is_active"; type:boolen;`
	Role     []Role    `gorm:"many2many:go_user_roles"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
