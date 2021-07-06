package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:create_time" json:"create_time"`
	UpdatedAt time.Time      `gorm:"column:modify_time" json:"modify_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
