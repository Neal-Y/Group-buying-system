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
	timezone := c.DefaultQuery("timezone", "UTC")
	params.PaginationParams = ParsePaginationParams(c)

	if startDateStr != "" && endDateStr != "" {
		params.StartDate, params.EndDate, err = ConvertDateRangeToUTC(startDateStr, endDateStr, timezone)
		if err != nil {
			return params, err
		}
	}

	return params, nil
}
