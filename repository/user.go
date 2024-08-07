package repository

import (
	"fmt"
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
	"time"
)

type UserRepository interface {
	Create(user *database.User) error
	FindByID(id int) (*database.User, error)
	Update(user *database.User) error
	FindByLineID(lineID string) (*database.User, error)
	SoftDeleteTx(tx *gorm.DB, id int) error
	BeginTransaction() *gorm.DB
	Upsert(user *database.User) error
	FindByDisplayName(displayName string) (*database.User, error)
	FindByEmailAndDisplayName(email, displayName string) (*database.User, error)
	SearchUsers(keyword string, startDate, endDate time.Time, offset int, limit int, isMember bool) ([]database.ExternalUser, int64, error)
	FindByIDAdmin(id int) (*database.ExternalUser, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: infrastructure.Db,
	}
}

func (r *userRepository) Create(user *database.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id int) (*database.User, error) {
	var user database.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *database.User) error {
	return r.db.Updates(user).Error
}

func (r *userRepository) FindByLineID(lineID string) (*database.User, error) {
	var user database.User
	err := r.db.Where("line_id = ?", lineID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) SoftDeleteTx(tx *gorm.DB, id int) error {
	var user database.User
	err := tx.First(&user, id).Error
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	user.DeletedAt = &now
	user.IsDeleted = true
	return tx.Save(&user).Error
}

func (r *userRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *userRepository) Upsert(user *database.User) error {
	existingUser, err := r.FindByLineID(user.LineID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if existingUser != nil {
		err = r.db.Model(existingUser).Updates(database.User{
			DisplayName: user.DisplayName,
			Email:       user.Email,
			LineToken:   user.LineToken,
			Phone:       user.Phone,
		}).Error
		if err != nil {
			fmt.Printf("Error updating user: %v\n", err)
		}
		return err
	} else {
		err = r.db.Create(user).Error
		if err != nil {
			fmt.Printf("Error creating user: %v\n", err)
		}
		return err
	}
}

func (r *userRepository) FindByDisplayName(displayName string) (*database.User, error) {
	var user database.User
	err := r.db.Where("display_name = ?", displayName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmailAndDisplayName(email, displayName string) (*database.User, error) {
	var user database.User
	err := r.db.Where("email = ? AND display_name = ? AND line_id = ?", email, displayName, "CreatedByUserEmail").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) SearchUsers(keyword string, startDate, endDate time.Time, offset int, limit int, isMember bool) ([]database.ExternalUser, int64, error) {
	var users []database.ExternalUser
	var count int64

	query := r.db.Model(&database.ExternalUser{})

	if keyword != "" {
		query = query.Where("id LIKE ? OR display_name LIKE ? OR email LIKE ? OR phone LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	if isMember != false {
		query = query.Where("is_member = ?", true)
	}

	err := query.Count(&count).
		Offset(offset).
		Limit(limit).
		Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (r *userRepository) FindByIDAdmin(id int) (*database.ExternalUser, error) {
	var user database.ExternalUser
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
