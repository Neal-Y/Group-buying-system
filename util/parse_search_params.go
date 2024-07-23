package util

import (
	"github.com/gin-gonic/gin"
	"time"
)

type SearchContainer struct {
	Keyword   string
	StartDate time.Time
	EndDate   time.Time
	PaginationParams
}

func SearchParams(c *gin.Context) (SearchContainer, error) {
	var params SearchContainer
	var err error

	params.Keyword = c.Query("keyword")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	params.PaginationParams = ParsePaginationParams(c)

	params.StartDate, err = ValidateTime(startDateStr)
	if err != nil {
		return params, err
	}

	params.EndDate, err = ValidateTime(endDateStr)
	if err != nil {
		return params, err
	}

	return params, nil
}
