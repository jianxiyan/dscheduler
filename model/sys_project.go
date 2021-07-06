package model

const projectPrefix = "dim_noah"

//项目表
type Project struct {
	MODEL
	Name        string `json:"name" gorm:"column:name;type:varchar(128);not null;unique;comment:名称"`
	IsDefault   int    `json:"is_default" gorm:"column:is_default;default:0;not null;comment:默认创建的个人项目"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
}

func (Project) TableName() string {
	return projectPrefix + "_project"
}

type ProjectPermission struct {
	MODEL
	Project         Project `json:"project" gorm:"foreignkey:Permissions;constraint:OnDelete:CASCADE;comment:项目"`
	AuthUser        User    `json:"auth_user" gorm:"foreignkey:ProjPermissions;constraint:OnDelete:CASCADE;comment:创建人"`
	Permissions     uint
	ProjPermissions uint
}

func (ProjectPermission) TableName() string {
	return projectPrefix + "_project_perimission"
}
