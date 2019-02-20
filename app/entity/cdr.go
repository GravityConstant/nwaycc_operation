package entity

type Cdr struct {
	Id               int64
	AccountId        int64  // 账号id
	Callee           string // 被叫号码
	Caller           string // 外显号码
	StartTime        string // 开始时间
	EndTime          string // 结束时间
	RouteId          int64  // 呼出路由
	Duration         int    // 通话时长
	BillBalance      float64
	RecordBase       string  // 呼叫录音目录-绝对路径，例如:/usr/local/record
	RecordPath       string  // 录音文件绝对地址，例如：base+/2018/03/xxxxx
	TaskId           int64   // 归属的任务id
	FeeRate          float64 // 归属的任务id
	CallId           string  // 呼叫唯一标识
	Intention        int     // 客户意向，0为肯定，1为否定，2为反感，3为意向
	Serverip         string  // 哪台服务器上的
	HangupDispostion string  // 哪方挂机：send_bye：平台挂机，recv_bye：被叫挂机，recv_cancel，呼叫被取消，recv_refuse:呼不通
	TermCause        string  // 挂机原因: 16:正常;非16:不正常
	TermStatus       string  // 挂机状态码：200：正常挂机，其他都不正常挂机
	TalkCrycle       int     // 对话轮数
	HasPushed        bool    // 是否已推送
}

// 用于查询的类
type CdrQueryParam struct {
	BaseQueryParam
	CallId      string
	StartTime   string
	EndTime     string
	DurationMin string
	DurationMax string
}

// cdr和task表联合查询
type CdrTask struct {
	Id               int64  `xlsx:"0"`
	Caller           string `xlsx:"1"`  // 外显号码
	StartTime        string `xlsx:"2"`  // 开始时间
	EndTime          string `xlsx:"3"`  // 结束时间
	Duration         int    `xlsx:"4"`  // 通话时长
	TaskId           int64  `xlsx:"5"`  // 归属的任务id
	TaskName         string `xlsx:"6"`  // 归属的任务名字
	CallId           string `xlsx:"7"`  // 呼叫唯一标识
	Intention        int    `xlsx:"8"`  // 客户意向，0为肯定，1为否定，2为反感，3为意向
	HangupDispostion string `xlsx:"9"`  // 哪方挂机：send_bye：平台挂机，recv_bye：被叫挂机，recv_cancel，呼叫被取消，recv_refuse:呼不通
	TermCause        string `xlsx:"10"` // 挂机原因: 16:正常;非16:不正常
	TermStatus       string `xlsx:"11"` // 挂机状态码：200：正常挂机，其他都不正常挂机
}
