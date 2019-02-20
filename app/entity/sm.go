package entity

import (
	"time"
)

type Sm struct {
	Id         int64
	TmpId      int64
	TaskId     int64
	CallNumber string
	CreateTime time.Time `orm:"null;type(datetime)"`
	HasSent    bool
	CallId     string
}

// 用于查询的类
type SmQueryParam struct {
	BaseQueryParam
	CallNumberLike string //模糊查询
	SearchHasSent  string //为空不查询，有值精确查询
}
