package service

import (
	"nway/nway_400/app/entity"
	"nway/nway_400/app/libs"
)

type cdrService struct{}

func (this *cdrService) table() string {
	return tableName("cdr")
}

func (this *cdrService) GetCdrList(params *entity.CdrQueryParam) ([]*entity.Cdr, int64) {
	data := make([]*entity.Cdr, 0)
	query := o.QueryTable(this.table())
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("account_id", entity.AccountId)
	if len(params.CallId) > 0 {
		query = query.Filter("call_id", params.CallId)
	} else {
		if len(params.StartTime) > 0 {
			query = query.Filter("start_time__gte", params.StartTime)
		}
		if len(params.EndTime) > 0 {
			query = query.Filter("end_time__lte", params.EndTime)
		}
		if len(params.DurationMin) > 0 {
			query = query.Filter("duration__gt", params.DurationMin)
		}
		if len(params.DurationMax) > 0 {
			query = query.Filter("duration__lt", params.DurationMax)
		}
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func (this *cdrService) GetTotal(params *entity.CdrQueryParam) (int64, error) {
	query := o.QueryTable(this.table())

	query = query.Filter("account_id", entity.AccountId)
	if len(params.CallId) > 0 {
		query = query.Filter("call_id", params.CallId)
	} else {
		if len(params.StartTime) > 0 {
			query = query.Filter("start_time__gte", params.StartTime)
		}
		if len(params.EndTime) > 0 {
			query = query.Filter("end_time__lte", params.EndTime)
		}
		if len(params.DurationMin) > 0 {
			query = query.Filter("duration__gt", params.DurationMin)
		}
		if len(params.DurationMax) > 0 {
			query = query.Filter("duration__lt", params.DurationMax)
		}
	}

	return query.Count()
}

func (this *cdrService) GetCdrTaskList(params *entity.CdrQueryParam) []*entity.CdrTask {
	data := make([]*entity.CdrTask, 0)
	query := new(libs.QueryString)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query.Filter("c", "account_id", entity.AccountId)
	if len(params.CallId) > 0 {
		query.Filter("c", "call_id", params.CallId)
	} else {
		if len(params.StartTime) > 0 {
			query.Filter("c", "start_time__gte", params.StartTime)
		}
		if len(params.EndTime) > 0 {
			query.Filter("c", "end_time__lte", params.EndTime)
		}
		if len(params.DurationMin) > 0 {
			query.Filter("c", "duration__gt", params.DurationMin)
		}
		if len(params.DurationMax) > 0 {
			query.Filter("c", "duration__lt", params.DurationMax)
		}
	}
	query.OrderBy("c", sortorder)
	query.Limit(params.Limit, params.Offset)

	var sql string
	sqlCol := `select c.id, c.caller, c.start_time, c.end_time, c.duration, c.task_id, t.tk_name task_name, c.call_id, c.intention, c.hangup_dispostion, c.term_cause, c.term_status `
	sqlFrom := `from ` + tableName("cdr") + ` c left join ` + tableName("task") + ` t on c.task_id=t.id `
	if len(query.FilterStr) > 0 {
		sql = sqlCol + sqlFrom + `where ` + query.String()
	} else {
		sql = sqlCol + sqlFrom + query.String()
	}
	o.Raw(sql).QueryRows(&data)

	return data
}
