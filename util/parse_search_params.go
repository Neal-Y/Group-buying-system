package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type SearchContainer struct {
	Keyword   string
	StartDate time.Time
	EndDate   time.Time
	Offset    int
	Limit     int
}

func ParseProductSearchParams(c *gin.Context) (SearchContainer, error) {
	var params SearchContainer
	var err error

	params.Keyword = c.Query("keyword")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	params.Offset, _ = strconv.Atoi(c.DefaultQuery("offset", "0"))
	params.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))

	if startDateStr != "" {
		params.StartDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return params, err
		}
	}
	if endDateStr != "" {
		params.EndDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return params, err
		}
	}

	return params, nil
}
