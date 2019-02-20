package service

import (
	"github.com/astaxie/beego/orm"
	"nway/nway_400/app/entity"
)

type smService struct{}

func (this *smService) table() string {
	return tableName("sm")
}

func (this *smService) GetSmList(params *entity.SmQueryParam) ([]*entity.Sm, int64) {
	data := make([]*entity.Sm, 0)
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

	query = query.Filter("callnumber__istartswith", params.CallNumberLike)
	if len(params.SearchHasSent) > 0 {
		query = query.Filter("HasSent", params.SearchHasSent)
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func (this *smService) GetTotal() (int64, error) {
	return o.QueryTable(this.table()).Count()
}

func (this *smService) UpdateColumns(column string, ids []int) (int64, error) {
	query := o.QueryTable(this.table())
	num, err := query.Filter("id__in", ids).Update(orm.Params{column: 1})
	return num, err
}
