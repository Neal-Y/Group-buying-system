package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer"
)

func (h *Post) updatePost(c *gin.Context) {
	id := c.Param("id")
	post := database.Post{}

	err := post.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	f := datatransfer.PostCreate{}
	err = c.ShouldBindJSON(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = post.Update(&f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}
