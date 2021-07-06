package model

const processorPrefix = "dim_noah_processor"

type Processor struct {
	MODEL
	Name string `json:"name" gorm:"column:name;type:varchar(128);not null;unique;comment:名称"`
	Code string `json:"code" gorm:"column:code;type:varchar(128);not null;unique;comment:唯一标识"`
	//1: python; 2: spark-sql; 3: spark-submit
	Type int8   `json:"type" gorm:"column:type;index;comment:插件类型"`
	Cmd  string `json:"cmd" gorm:"column:cmd;type:text;comment:插件执行命令"`
	//1: 共享; 2: 私有
	PublicType  int8   `json:"public_type" gorm:"column:public_type;index;comment:开放类型"`
	Config      string `json:"config" gorm:"column:config;type:text;comment:插件执行命令"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
}

func (Processor) TableName() string {
	return processorPrefix + "_project"
}

type ProcessorVersion struct {
	MODEL
	Processor Processor `json:"processor" gorm:"foreignkey:PVersion;constraint:OnDelete:CASCADE;comment:插件"`
	PVersion  uint

	Name string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:名称"`
	//插件使用状态： 0没用 1使用中
	Status      int8   `json:"status" gorm:"column:status;default:0;comment:使用状态"`
	Version     string `json:"version" gorm:"column:version;type:varchar(128);comment:版本"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
}

func (ProcessorVersion) TableName() string {
	return processorPrefix + "_version"
}
