package model

import (
	"time"

	"gorm.io/gorm"
)

// constraint:OnDelete:
// CASCADE删除关联数据的时候，与之的关联也删除
// DO_NOTHING删除关联数据的时候，什么操作也不做
// PROTRCT删除关联数据的时候，引发报错
// SET_NULL删除关联数据的时候，与之关联的只设置为空
// SET_DEFAULT删除关联数据的时候，与之关联的只设置为默认值
// SET删除关联数据

type MODELBASE struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:create_time" json:"create_time"`
	UpdatedAt time.Time      `gorm:"column:modify_time" json:"modify_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type MODEL struct {
	MODELBASE
	CreateUser      User `json:"create_user" gorm:"foreignkey:CreateUserRefer;constraint:OnDelete:SET NULL;comment:创建人"`
	ModifyUser      User `json:"modify_user" gorm:"foreignkey:ModifyUserRefer;constraint:OnDelete:SET NULL;comment:更新人"`
	CreateUserRefer uint
	ModifyUserRefer uint
}
