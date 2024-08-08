package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"shopping-cart/constant"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/admin"
	"shopping-cart/repository"
	"shopping-cart/util"
)

type AdminService interface {
	RegisterAdmin(admin *admin.Request) error
	Login(req *admin.Login) (string, error)
	GetAdmin() (*database.Admin, error)
	GetAdminByUsername(username string) (*database.Admin, error)
	UpdateAdmin(id int, req *admin.UpdateRequest) (*database.Admin, error)
	RequestPasswordReset(email string) error
	ResetPassword(email, code, newPassword string) error
}

type adminService struct {
	adminRepo  repository.AdminRepository
	verifyRepo repository.VerifyRepository
}

func NewAdminService(adminRepo repository.AdminRepository, verifyRepo repository.VerifyRepository) AdminService {
	return &adminService{
		adminRepo:  adminRepo,
		verifyRepo: verifyRepo,
	}
}

func (s *adminService) RegisterAdmin(req *admin.Request) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin := &database.Admin{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Email:        req.Email,
		LineID:       req.LineID,
	}
	return s.adminRepo.Create(admin)
}

func (s *adminService) Login(req *admin.Login) (string, error) {
	username, password := req.Username, req.Password

	admin, err := s.adminRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := util.GenerateJWT(constant.AdminType)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *adminService) GetAdminByUsername(username string) (*database.Admin, error) {
	return s.adminRepo.FindByUsername(username)
}

func (s *adminService) GetAdmin() (*database.Admin, error) {
	return s.adminRepo.GetAdmin()
}

func (s *adminService) UpdateAdmin(id int, req *admin.UpdateRequest) (*database.Admin, error) {
	admin, err := s.adminRepo.GetAdmin()
	if err != nil {
		return nil, err
	}

	if req.Username != "" {
		admin.Username = req.Username
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		admin.PasswordHash = string(hashedPassword)
	}
	if req.Email != "" {
		admin.Email = req.Email
	}

	if req.LineID != "" {
		admin.LineID = req.LineID
	}

	err = s.adminRepo.Update(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *adminService) RequestPasswordReset(email string) error {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	err := s.verifyRepo.SaveVerificationCode(email, code)
	if err != nil {
		return err
	}

	err = util.SendResetCodeEmail(email, code)
	if err != nil {
		return err
	}

	return nil
}

func (s *adminService) ResetPassword(email, code, newPassword string) error {
	savedCode, err := s.verifyRepo.GetVerificationCode(email)
	if err != nil || savedCode != code {
		return errors.New("invalid or expired verification code or code has already been used")
	}

	err = s.verifyRepo.MarkCodeAsUsed(email)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin, err := s.adminRepo.FindByEmail(email)
	if err != nil {
		return errors.New("admin not found")
	}
	admin.PasswordHash = string(hashedPassword)
	return s.adminRepo.Update(admin)
}
