package repository

import (
	"github.com/tsfans/go/framework"
	"github.com/tsfans/go/framework/database"
	"github.com/tsfans/go/server/model/dto"
)

func QueryDataById[T any](id uint) (data *T, err error) {
	data, err = QueryData[T]("id = ?", id)
	return
}

func QueryData[T any](query any, args ...any) (data *T, err error) {
	data = new(T)
	rpn := database.DB.Where(query, args).Find(&data)
	if rpn.Error != nil {
		err = framework.NewServiceError(dto.SC_DB_ERROR, rpn.Error.Error())
	}
	if rpn.RowsAffected == 0 {
		data = nil
	}
	return
}

func QueryDatas[T any](query any, args ...any) (datas []T, err error) {
	rpn := database.DB.Where(query, args).Find(&datas)
	if rpn.Error != nil {
		err = framework.NewServiceError(dto.SC_DB_ERROR, rpn.Error.Error())
	}
	return
}

func PageQueryDatas[T any](pageReq *dto.PageRequest, query any, args ...any) (datas []T, total int64, err error) {
	table := new(T)
	filter := database.DB.Model(&table)
	if query != nil {
		filter = filter.Where(query, args)
	}
	rpn := filter.Count(&total)
	if rpn.Error != nil {
		err = framework.NewServiceError(dto.SC_DB_ERROR, rpn.Error.Error())
	}
	if total == 0 {
		return
	}
	rpn = filter.Offset(pageReq.GetOffset()).Limit(pageReq.GetPageSize()).Order(pageReq.GetSort()).Find(&datas)
	if rpn.Error != nil {
		err = framework.NewServiceError(dto.SC_DB_ERROR, rpn.Error.Error())
	}
	return
}

func CreateData(data any) error {
	rpn := database.DB.Create(data)
	if rpn.Error != nil {
		return framework.NewServiceError(dto.SC_DB_ERROR, rpn.Error.Error())
	}
	return nil
}

func UpdateDataById[T any](update map[string]any, id uint) error {
	return UpdateData[T](update, "id = ?", id)
}

func UpdateData[T any](update map[string]any, query any, args ...any) error {
	table := new(T)
	rpn := database.DB.Model(&table).Where(query, args).Updates(update)
	if rpn.Error != nil {
		return framework.NewServiceError(dto.SC_DB_ERROR, rpn.Error.Error())
	}
	return nil
}

func DeleteById[T any](id uint) error {
	return Delete[T]("id = ?", id)
}

func Delete[T any](query any, args ...any) error {
	table := new(T)
	rpn := database.DB.Model(&table).Delete(query, args)
	if rpn.Error != nil {
		return framework.NewServiceError(dto.SC_DB_ERROR, rpn.Error.Error())
	}
	return nil
}
