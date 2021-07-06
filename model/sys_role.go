package model

const rolePrefix = "dim_ark"

//角色信息
type Role struct {
	MODEL
	Code        string `json:"code" gorm:"column:code;type:varchar(128);not null;unique;comment:角色标识"`
	Src         string `json:"src" gorm:"column:src;type:varchar(32);comment:来源"`
	Name        string `json:"name" gorm:"column:name;type:varchar(128);not null;unique;comment:角色名"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
}

func (Role) TableName() string {
	return rolePrefix + "_role"
}

//角色信息
type UserRole struct {
	MODEL
	Role        Role `json:"role" gorm:"foreignkey:UserRoles;constraint:OnDelete:CASCADE;comment:角色"`
	User        User `json:"auth_user" gorm:"foreignkey:RoleUser;constraint:OnDelete:CASCADE;comment:创建人"`
	UserRoles   uint
	RoleUser    uint
	Src         string `json:"src" gorm:"column:src;type:varchar(32);comment:来源"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
}

func (UserRole) TableName() string {
	return rolePrefix + "user_role"
}

//角色-父子关系
type RoleParent struct {
	MODEL
	Role        Role `json:"project" gorm:"foreignkey:RoleParent;constraint:OnDelete:CASCADE;comment:父规则"`
	Parent      Role `json:"parent" gorm:"foreignkey:RoleChild;constraint:OnDelete:CASCADE;comment:子规则"`
	RoleParent  uint
	RoleChild   uint
	Src         string `json:"src" gorm:"column:src;type:varchar(32);comment:来源"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
}

func (RoleParent) TableName() string {
	return rolePrefix + "role_parent"
}
