package utils

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Pagination struct {
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	TotalData int64       `json:"totalData"`
	TotalPage int         `json:"totalPage"`
	Data      interface{} `json:"data"`
}

func Paginate[T any](c *fiber.Ctx, db *gorm.DB, out *[]T) (*Pagination, error) {
	limit, _ := strconv.Atoi(c.Params("limit", "10"))
	page, _ := strconv.Atoi(c.Params("page", "0"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var totalData int64
	if err := db.Model(out).Count(&totalData).Error; err != nil {
		return nil, err
	}

	if err := db.Offset(offset).Limit(limit).Find(out).Error; err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	return &Pagination{
		Page:      page,
		Limit:     limit,
		TotalData: totalData,
		TotalPage: totalPage,
		Data:      out,
	}, nil
}
