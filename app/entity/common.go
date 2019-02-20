package entity

// BaseQueryParam 用于查询的类
type BaseQueryParam struct {
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Offset int64  `json:"offset"`
	Limit  int    `json:"limit"`
}

// JsonResult 用于返回ajax请求的基类
type JsonResult struct {
	Code JsonResultCode `json:"code"`
	Msg  string         `json:"msg"`
	Obj  interface{}    `json:"obj"`
}

type JsonResultCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问
)

const BaseFormat = "2006-01-02 15:04:05"

// 中企的accountid默认是151
const AccountId = 151
