package post

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"shopping-cart/model/database"
)

func (h *Post) get(c *gin.Context) {
	posts := database.Posts{}

	err := posts.FindAll()
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Post) getByID(c *gin.Context) {
	id := c.Param("id")
	post := database.Post{}

	err := post.FindById(id).Error

	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, post)
}
