package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type PaginationParams struct {
	Offset int
	Limit  int
}

func ParsePaginationParams(c *gin.Context) PaginationParams {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	return PaginationParams{
		Offset: offset,
		Limit:  limit,
	}
}
