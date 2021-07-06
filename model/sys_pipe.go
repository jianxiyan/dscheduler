package model

import "time"

const pipePrefix = "dim_noah"

//流程
type Pipeline struct {
	MODEL
	Name     string  `json:"name" gorm:"column:name;type:varchar(128);not null;unique;comment:名称"`
	Project  Project `json:"project" gorm:"foreignkey:PipeLine;constraint:OnDelete:CASCADE;comment:项目"`
	PipeLine uint
	CtTime   string `json:"ct_time" gorm:"column:ct_time;type:varchar(128);comment:调度时间"`
	//1: 已下线; 2: 已上线; 3: 已过期
	Status int8 `json:"status" gorm:"column:status;default:1;comment:上线状态"`
	//0: 平台; 1: API
	Src       int8   `json:"src" gorm:"column:src;comment:来源"`
	Tag       string `json:"tag" gorm:"column:tag;type:varchar(1024);default:;comment:标签"`
	LifeCycle string `json:"life_cycle" gorm:"column:life_cycle;type:varchar(50);comment:流程过期时"`
	//2: 钉钉; 3: 邮件; 4: 短信
	MonitorWay string `json:"monitor_way" gorm:"column:monitor_way;type:varchar(50);comment:报警方式"`
	//1: 共享; 2: 私有
	PublicType  int8   `json:"public_type" gorm:"column:public_type;index;default:1;comment:开放类型"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`

	CallbackUri string `json:"callback_uri" gorm:"column:callback_uri;type:varchar(256);comment:回调url"`
	Config      string `json:"config" gorm:"column:config;type:text;comment:参数配置项"`
}

func (Pipeline) TableName() string {
	return pipePrefix + "_pipeline"
}

//任务
type Task struct {
	MODEL
	Pipeline       Pipeline `json:"pipeline" gorm:"foreignkey:PipelineTasks;constraint:OnDelete:CASCADE;comment:流程"`
	PipelineTasks  uint
	Processor      Processor `json:"processor" gorm:"foreignkey:ProcessorTasks;constraint:OnDelete:CASCADE;comment:流程"`
	ProcessorTasks uint

	Name string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:名称"`
	//1: 已下线; 2: 已上线;
	Status      int8   `json:"status" gorm:"column:status;default:1;comment:上线状态"`
	NextTaskIds string `json:"next_task_ids" gorm:"column:next_task_ids;type:varchar(1024);default:,;comment:下游任务id列表"`
	PrevTaskIds string `json:"prev_task_ids" gorm:"column:prev_task_ids;type:varchar(1024);default:,;comment:上游任务id列表"`
	Config      string `json:"config" gorm:"column:config;type:text;comment:参数配置项"`
	Index       string `json:"index" gorm:"column:index;type:text;comment:任务坐标"`
	RetryCount  int8   `json:"retry_count" gorm:"column:retry_count;default:0;comment:任务失败重试次数"`
	LastRunTime string `json:"last_run_time" gorm:"column:last_run_time;type:varchar(50);comment:最后执行时间"`
	Priority    int8   `json:"priority" gorm:"column:priority;default:6;comment:优先级"`
	// //0: 否; 1: 是;
	// IsFailedTsouch      int8      `json:"is_failed_touch" gorm:"column:is_failed_touch;default:0;comment:是否失败触发下游"`
	ServerTag   string `json:"server_tag" gorm:"column:server_tag;type:varchar(50);default:ALL;comment:运行服务器"`
	Description string `json:"description" gorm:"column:description;type:text;comment:描述"`
}

func (Task) TableName() string {
	return pipePrefix + "_task"
}

//任务
type RunHistory struct {
	MODELBASE
	Pipeline    Pipeline `json:"pipe_history" gorm:"foreignkey:PipeHistory;constraint:OnDelete:CASCADE;comment:流程"`
	PipeHistory uint
	Task        Task `json:"task" gorm:"foreignkey:TaskHistory;constraint:OnDelete:CASCADE;comment:任务"`
	TaskHistory uint

	RunTime   string    `json:"run_time" gorm:"column:run_time;type:varchar(20);index;comment:执行时间"`
	StartTime time.Time `json:"start_time" gorm:"column:start_time;default:2000-01-01 00:00:00;comment:开始时间"`
	EndTime   time.Time `json:"end_time" gorm:"column:end_time;default:2000-01-01 00:00:00;comment:结束时间"`
	//0: 等待中; 1: 正在执行; 2: 执行成功; 3: 执行失败; 4: 任务超时; 5: 等待调度; 6: 任务被用户停止; 7: 上游任务失败;
	Status    int8 `json:"status" gorm:"column:status;default:1;index;comment:运行状态"`
	Triggered int8 `json:"triggered" gorm:"column:triggered;default:0;index;comment:是否触发"`

	//调度类型, 1:并行, 2.串行
	ScheduleType int8 `json:"schedule_type" gorm:"column:schedule_type;default:1;index;comment:调度类型"`
	Priority     int8 `json:"priority" gorm:"column:priority;default:6;comment:优先级"`
	//任务来源，1.每天定时执行，2：手动添加，3:拓扑唤醒,4:失败重新提交,10:未知
	TaskFrom  int8 `json:"task_from" gorm:"column:task_from;default:1;comment:任务来源"`
	RetryTime int8 `json:"retry_time" gorm:"column:retry_time;default:0;comment:失败重试次数"`

	TaskHandler string `json:"task_handler" gorm:"column:task_handler;type:varchar(1024);comment:任务执行句抦，一般为进程id"`
	Hostname    string `json:"hostname" gorm:"column:hostname;type:varchar(256);comment:执行任务机器"`
	ServerTag   string `json:"server_tag" gorm:"column:server_tag;type:varchar(50);default:ALL;comment:可运行服务器"`
}

func (RunHistory) TableName() string {
	return pipePrefix + "_run_history"
}

type PipelinePermission struct {
	MODEL
	Pipeline        Pipeline `json:"pipeline" gorm:"foreignkey:Permissions;constraint:OnDelete:CASCADE;comment:流程"`
	Permissions     uint
	AuthUser        User `json:"auth_user" gorm:"foreignkey:PipePermissions;constraint:OnDelete:CASCADE;comment:创建人"`
	PipePermissions uint
}

func (PipelinePermission) TableName() string {
	return processorPrefix + "_pipe_permission"
}

type PipelineAlarm struct {
	MODEL
	Pipeline          Pipeline `json:"pipeline" gorm:"foreignkey:AlarmPipe;constraint:OnDelete:CASCADE;comment:流程"`
	AlarmPipe         uint
	User              User `json:"user" gorm:"foreignkey:PipelineAlarmUser;constraint:OnDelete:CASCADE;comment:创建人"`
	PipelineAlarmUser uint
}

func (PipelineAlarm) TableName() string {
	return processorPrefix + "_pipe_alarm"
}
