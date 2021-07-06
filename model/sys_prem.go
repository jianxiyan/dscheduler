package model

const premPrefix = "dim_ark"

//角色信息
type SynData struct {
	MODEL
	Name string `json:"name" gorm:"column:name;type:varchar(128);not null;unique;comment:角色名"`
	//1: 控制； 2: 不控制
	PermFlag    uint8  `json:"prem_flag" gorm:"column:prem_flag;default:2;comment:是否控制权限"`
	Sql         string `json:"sql" gorm:"column:sql;type:text;comment:同步数据源"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
}

func (SynData) TableName() string {
	return premPrefix + "_syn_data"
}

// //数据权限表-角色
// type DataAuth struct {
// 	MODEL
// 	Role        Role `json:"role" gorm:"foreignkey:RoleDatas;constraint:OnDelete:CASCADE;comment:角色"`
// 	ContentType        Count `json:"content_type" gorm:"foreignkey:RoleUser;constraint:OnDelete:CASCADE;comment:创建人"`
// 	RoleDatas   uint
// 	RoleUser    uint
// 	Src         string `json:"src" gorm:"column:src;type:varchar(32);comment:来源"`
// 	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
// }

// func (DataAuth) TableName() string {
// 	return rolePrefix + "data_auth"
// }
