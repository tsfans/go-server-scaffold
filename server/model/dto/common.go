package dto

import (
	"time"

	"github.com/tsfans/go/framework"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (res *Response) Success(data any) {
	res.Code = SC_SUCCESS
	res.Msg = "success"
	res.Data = data
}

func (res *Response) Fail(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

func (res *Response) FailWithError(err error) {
	if serr, ok := err.(*framework.ServiceError); ok {
		res.Fail(serr.Code, serr.Msg)
	} else {
		res.Fail(SC_INTERNAL_ERROR, err.Error())
	}
}

type PageRequest struct {
	PageNum   *int    `json:"pageNum" form:"pageNum"`
	PageSize  *int    `json:"pageSize" form:"pageSize"`
	SearchKey *string `json:"searchKey" form:"searchKey"`
	Sort      *string `json:"sort" form:"sort"`
}

func (p PageRequest) GetPageNum() int {
	if p.PageNum != nil {
		pageNum := *p.PageNum
		if pageNum < 1 {
			return 1
		}
		return pageNum
	}
	return 1
}

func (p PageRequest) GetPageSize() int {
	if p.PageSize != nil {
		pageSize := *p.PageSize
		if pageSize < 1 {
			return 10
		}
		return pageSize
	}
	return 10
}

func (p PageRequest) GetSort() string {
	if p.Sort != nil {
		return *p.Sort
	}
	return "id"
}

func (p PageRequest) GetOffset() int {
	return p.GetPageSize() * (p.GetPageNum() - 1)
}

type Page struct {
	PageNum   int   `json:"pageNum"`
	PageSize  int   `json:"pageSize"`
	TotalPage int   `json:"totalPage"`
	Total     int64 `json:"total"`
	Datas     []any `json:"datas"`
}

func NewPage(pageNum int, pageSize int, total int64, datas []any) *Page {
	page := &Page{
		PageNum:   pageNum,
		PageSize:  pageSize,
		TotalPage: 0,
		Total:     total,
		Datas:     datas,
	}
	if pageSize == 0 {
		return page
	}

	var totalPage int64
	n := total % int64(pageSize)
	if n == 0 {
		totalPage = total / int64(pageSize)
	} else {
		totalPage = total/int64(pageSize) + 1
	}
	page.TotalPage = int(totalPage)

	return page
}

type Common struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UriId struct {
	Id uint `uri:"id" form:"id" binding:"required"`
}
