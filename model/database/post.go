package database

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/datatransfer"
	"time"
)

// Post model
type Post struct {
	Id        int       `json:"id",gorm:"primarykey"`
	Title     string    `json:"title",gorm:"column:title"`
	Content   string    `json:"content",gorm:"column:content"`
	CreatedAt time.Time `json:"created_at",gorm:"column:created_at"`
}

func (Post) TableName() string {
	return "posts"
}

func (post *Post) model() *gorm.DB { return infrastructure.Db.Model(post) }

func (post *Post) FindAll() error {
	return post.model().Find(&post).Error
}

func (post *Post) FindById(id string) error {
	return post.model().Find(&post, id).Error
}

func (post *Post) Update(updatePayload *datatransfer.PostCreate) error {
	updateData := Post{
		Title:   updatePayload.Title,
		Content: updatePayload.Content,
	}

	return post.model().Updates(updateData).Error
}

func (post *Post) Delete() error {
	return post.model().Delete(post).Error
}

func (post *Post) Create() error {
	return post.model().Create(post).Error
}

type Posts []Post

func (posts *Posts) model() *gorm.DB { return infrastructure.Db.Model(posts) }

func (posts *Posts) FindAll() error {
	return posts.model().Find(&posts).Error
}
