package model

const menuPrefix = "dim_analysis"

type Menu struct {
	MODEL
	Name     string `json:"name,omitempty" gorm:"column:name;type:varchar(64);not null;comment:姓名"`
	Parentid *Menu  `json:"parentid" gorm:"foreignkey:Self;constraint:OnDelete:CASCADE;comment:项目"`
	Self     int32

	PrevMenuId string `json:"prev_menu_id,omitempty" gorm:"column:prev_menu_id;type:varchar(64);default:;comment:上游菜单id"`
	NextMenuId string `json:"next_menu_id,omitempty" gorm:"column:next_menu_id;type:varchar(64);default:;comment:下游菜单id"`
	Path       string `json:"path,omitempty" gorm:"column:path;type:varchar(128);not null;unique;comment:菜单路由"`
	Link       string `json:"link,omitempty" gorm:"column:link;type:varchar(128);comment:链接"`
	Level      string `json:"level,omitempty" gorm:"column:level;type:varchar(32);comment:菜单等级"`
	Icon       string `json:"icon,omitempty" gorm:"column:icon;type:varchar(32);comment:图标"`
	Component  string `json:"component,omitempty" gorm:"column:component;type:varchar(64);comment:组件"`
	Order      string `json:"order,omitempty" gorm:"column:order;type:varchar(32);not null;comment:顺序"`
	//1: 下线； 2: 上线
	Status bool `json:"status,omitempty" gorm:"column:status;default:1;comment:上线状态"`
	//1: 目录； 2:页面
	Type bool `json:"type,omitempty" gorm:"column:type;default:1;comment:菜单种类"`

	PlatformId  string `json:"platform_id,omitempty" gorm:"column:platform_id;type:varchar(64);comment:平台唯一标识"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
}

func (Menu) TableName() string {
	return menuPrefix + "_menu"
}
