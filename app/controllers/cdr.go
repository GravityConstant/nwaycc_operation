package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"nway/nway_400/app/entity"
	"nway/nway_400/app/libs"
	"nway/nway_400/app/service"

	"github.com/tealeg/xlsx"
)

type CdrController struct {
	BaseController
}

// 首页
func (this *CdrController) Index() {
	this.Data["pageTitle"] = "话单详情"

	this.display()
}

func (this *CdrController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params entity.CdrQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)

	data := service.CdrService.GetCdrTaskList(&params)
	total, _ := service.CdrService.GetTotal(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	this.Data["json"] = result
	this.ServeJSON()
}

//golang 获取数据库数据导出excel 存入文件
func (this *CdrController) DownloadCdr() {
	// 接收参数
	StartTime := strings.TrimSpace(this.GetString("StartTime"))
	EndTime := strings.TrimSpace(this.GetString("EndTime"))
	CallId := strings.TrimSpace(this.GetString("CallId"))
	DurationMin := strings.TrimSpace(this.GetString("DurationMin"))
	DurationMax := strings.TrimSpace(this.GetString("DurationMax"))

	var params entity.CdrQueryParam
	params.CallId = CallId
	params.StartTime = StartTime
	params.EndTime = EndTime
	params.DurationMin = DurationMin
	params.DurationMax = DurationMax
	params.Limit = 3000

	total, err := service.CdrService.GetTotal(&params)
	if err == nil && total > 3000 {
		this.DownloadCdrUnLimit(&params, total)
		return
	}

	fmt.Println("CdrController, DownloadCdr==========================")
	fmt.Printf("%#v\n", params)

	data := service.CdrService.GetCdrTaskList(&params)

	var rows [][]string
	th := []string{"ID", "主叫号码", "开始时间", "结束时间", "通话时长", "任务id", "任务归属", "呼叫标识", "客户意向", "挂机方向", "挂机原因", "挂机状态码"}
	rows = append(rows, th)
	var tt string
	// fmt.Println("=======================")
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		tr := make([]string, 0)
		tr = append(tr, strconv.FormatInt(data[i].Id, 10))
		tr = append(tr, data[i].Caller)
		startTime := strings.Split(data[i].StartTime, " ")
		if len(startTime) > 1 {
			tt = startTime[0] + " " + startTime[1]
		} else {
			tt = ""
		}
		tr = append(tr, tt)
		endTime := strings.Split(data[i].EndTime, " ")
		if len(startTime) > 1 {
			tt = endTime[0] + " " + endTime[1]
		} else {
			tt = ""
		}
		tr = append(tr, tt)
		tr = append(tr, strconv.Itoa(data[i].Duration))
		tr = append(tr, strconv.FormatInt(data[i].TaskId, 10))
		tr = append(tr, data[i].TaskName)
		tr = append(tr, data[i].CallId)
		tr = append(tr, strconv.Itoa(data[i].Intention))
		tr = append(tr, data[i].HangupDispostion)
		tr = append(tr, data[i].TermCause)
		tr = append(tr, data[i].TermStatus)
		rows = append(rows, tr)
	}

	if filePath, err := libs.ExportExcel(rows, "aicdr.xlsx"); err != nil {
		fmt.Printf(err.Error())
	} else {
		this.Ctx.Output.Download(filePath, "aicdr.xlsx")
	}
}

func (this *CdrController) DownloadCdrUnLimit(params *entity.CdrQueryParam, total int64) {
	params.Limit = 5000
	data := service.CdrService.GetCdrTaskList(params)

	f := xlsx.NewFile()
	sheet, _ := f.AddSheet("aicdr1")
	th := sheet.AddRow()
	thCon := struct {
		Id               string
		Caller           string
		StartTime        string
		EndTime          string
		Duration         string
		TaskId           string
		TaskName         string
		CallId           string
		Intention        string
		HangupDispostion string
		TermCause        string
		TermStatus       string
	}{
		Id:               "ID",
		Caller:           "主叫号码",
		StartTime:        "开始时间",
		EndTime:          "结束时间",
		Duration:         "通话时长",
		TaskId:           "任务id",
		TaskName:         "任务归属",
		CallId:           "呼叫标识",
		Intention:        "客户意向",
		HangupDispostion: "挂机方向",
		TermCause:        "挂机原因",
		TermStatus:       "挂机状态码",
	}
	th.WriteStruct(&thCon, -1)
	for _, trCon := range data {
		tr := sheet.AddRow()
		tr.WriteStruct(trCon, -1)
	}

	// 判断是否需要重新获取数据
	dataLen := int64(len(data))
	if total > dataLen {
		counts := int(total / dataLen)

		for i := 0; i < counts; i++ {
			params.Offset = int64((i + 1) * params.Limit)
			data = service.CdrService.GetCdrTaskList(params)
			go func() {
				sheet, _ = f.AddSheet("aicdr" + strconv.Itoa(i+2))
				th = sheet.AddRow()
				th.WriteStruct(&thCon, -1)
				for _, trCon := range data {
					tr := sheet.AddRow()
					tr.WriteStruct(trCon, -1)
				}
			}()
		}
	}
	filePath := "logs/aicdr_long.xlsx"
	if err := f.Save(filePath); err == nil {
		this.Ctx.Output.Download(filePath, "aicdr_long.xlsx")
	} else {
		fmt.Printf(err.Error())
	}
}
