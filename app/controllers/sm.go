package controllers

import (
	"encoding/json"
	"fmt"
	// "github.com/astaxie/beego"
	"nway/nway_400/app/entity"
	"nway/nway_400/app/libs"
	"nway/nway_400/app/service"
	"strconv"
	"strings"
)

type SmController struct {
	BaseController
}

// 首页
func (this *SmController) Index() {
	this.Data["pageTitle"] = "短信查询"

	this.display()
}

func (this *SmController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params entity.SmQueryParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)

	data, total := service.SmService.GetSmList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *SmController) UpdateHasSent() {
	strs := this.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := service.SmService.UpdateColumns("HasSent", ids); err == nil {
		this.jsonResult(entity.JRCodeSucc, fmt.Sprintf("成功更新 %d 项", num), 0)
	} else {
		this.jsonResult(entity.JRCodeFailed, "更新失败", 0)
	}
}

//golang 获取数据库数据导出excel 存入文件
func (this *SmController) DownloadSm() {
	var params entity.SmQueryParam
	params.Order = "desc"
	data, _ := service.SmService.GetSmList(&params)

	lastId := libs.ReadExcelLastId("logs/aism.xlsx")
	if len(lastId) > 0 {
		lastId2Int, _ := strconv.ParseInt(lastId, 10, 64)

		if !(data[0].Id > lastId2Int) {
			this.Ctx.Output.Download("logs/aism.xlsx", "aism.xlsx")
			return
		}
	}

	var rows [][]string
	th := []string{"ID", "话术Id", "任务Id", "呼叫号码", "创建时间", "是否已发送短信", "呼叫标识"}
	rows = append(rows, th)
	// fmt.Println("=======================")
	for i := 0; i < len(data); i++ {
		tr := []string{}
		tr = append(tr, strconv.FormatInt(data[i].Id, 10))
		tr = append(tr, strconv.FormatInt(data[i].TmpId, 10))
		tr = append(tr, strconv.FormatInt(data[i].TaskId, 10))
		tr = append(tr, data[i].CallNumber)
		tr = append(tr, data[i].CreateTime.Format(entity.BaseFormat))
		tr = append(tr, strconv.FormatBool(data[i].HasSent))
		tr = append(tr, data[i].CallId)
		rows = append(rows, tr)
	}
	filePath, _ := libs.ExportExcel(rows, "aism.xlsx")
	this.Ctx.Output.Download(filePath, "aism.xlsx")
}
