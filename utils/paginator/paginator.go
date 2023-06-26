package paginator

import (
	"math"

	"gorm.io/gorm"
)

type Meta struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int   `json:"total_page"`
	Page        int   `json:"page"`
	Prev        *int  `json:"prev"`
	Next        *int  `json:"next"`
}

func Paginator(page, limit int, total int64) (meta Meta) {

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	meta.Page = page
	meta.TotalPage = int(math.Ceil(float64(total) / float64(limit)))
	meta.TotalRecord = total

	var next, prev int

	if page >= meta.TotalPage {

	} else {
		next = page + 1
		meta.Next = &next
	}

	if page > 1 {
		prev = page - 1
		meta.Prev = &prev
	} else {
	}
	return meta

}

func Paginate(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
	  	offset := (page - 1) * limit
	  	return db.Offset(offset).Limit(limit)
	}
}