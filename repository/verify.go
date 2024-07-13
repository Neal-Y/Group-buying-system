package repository

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
)

type VerifyRepository interface {
	SaveVerificationCode(email, code string) error
	GetVerificationCode(email string) (string, error)
	MarkCodeAsUsed(email string) error
}

type verifyRepository struct {
	db *gorm.DB
}

func NewVerifyRepository() VerifyRepository {
	return &verifyRepository{
		db: infrastructure.Db,
	}
}

func (r *verifyRepository) SaveVerificationCode(email, code string) error {
	verificationCode := &database.VerificationCode{
		Email: email,
		Code:  code,
		Used:  false,
	}
	return r.db.Create(verificationCode).Error
}

func (r *verifyRepository) GetVerificationCode(email string) (string, error) {
	var verificationCode database.VerificationCode
	err := r.db.Where("email = ? AND used = ?", email, false).Order("created_at desc").First(&verificationCode).Error
	if err != nil {
		return "", err
	}
	return verificationCode.Code, nil
}

func (r *verifyRepository) MarkCodeAsUsed(email string) error {
	return r.db.Model(&database.VerificationCode{}).Where("email = ? AND used = ?", email, false).Update("used", true).Error
}
