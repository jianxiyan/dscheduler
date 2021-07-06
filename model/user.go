package model

const userPrefix = "dim_ark"

type User struct {
	MODELBASE
	UserId   string `json:"userid,omitempty" gorm:"column:userid;type:varchar(32);not null;unique;comment:用户标识"`
	Name     string `json:"name,omitempty" gorm:"column:name;type:varchar(32);not null;comment:姓名"`
	UserName string `json:"user_name,omitempty" gorm:"column:user_name;type:varchar(32);not null;unique;comment:用户名"`
	Password string `json:"password,omitempty" gorm:"column:password;type:varchar(100);not null;comment:用户登录密码"`
	Src      string `json:"src,omitempty" gorm:"column:src;type:varchar(32);comment:来源"`
	NickName string `json:"nick_name,omitempty" gorm:"column:nick_name;type:varchar(32);comment:昵称"`
	Mobile   string `json:"mobile,omitempty" gorm:"column:mobile;type:varchar(32);comment:手机号码"`
	Email    string `json:"email,omitempty" gorm:"column:email;type:varchar(128);comment:邮箱"`
	ApiPwd   string `json:"api_pwd,omitempty" gorm:"column:api_pwd;type:varchar(128);comment:api验证码"`
	IsActive bool   `json:"is_active,omitempty" gorm:"column:is_active;default:1;comment:api验证码"`
	DingKey  string `json:"ding_key,omitempty" gorm:"column:ding_key;type:varchar(32);comment:用户钉钉标识"`
}

func (User) TableName() string {
	return userPrefix + "_user"
}
