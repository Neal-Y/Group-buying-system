package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/database"
)

func (h *Post) deletePost(c *gin.Context) {
	id := c.Param("id")
	post := database.Post{}

	err := post.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = post.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
