package service

import (
	"errors"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/admin"
	"shopping-cart/model/datatransfer/user"
	"shopping-cart/repository"
	"shopping-cart/util"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	RegisterAdmin(admin *admin.Request) error
	Login(username, password string) (string, error)
	GetAdminByID(id int) (*database.Admin, error)
	UpdateAdmin(id int, req *admin.UpdateRequest) (*database.Admin, error)
	DeleteAdmin(id int) error
	CreateUser(user *user.Request) error
	GetUserByID(id int) (*database.User, error)
	UpdateUser(id int, user *user.Request) error
	DeleteUser(id int) error
}

type adminService struct {
	adminRepo repository.AdminRepository
	userRepo  repository.UserRepository
	orderRepo repository.OrderRepository
}

func NewAdminService(adminRepo repository.AdminRepository, userRepo repository.UserRepository, orderRepo repository.OrderRepository) AdminService {
	return &adminService{
		adminRepo: adminRepo,
		userRepo:  userRepo,
		orderRepo: orderRepo,
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
	}
	return s.adminRepo.Create(admin)
}

func (s *adminService) Login(username, password string) (string, error) {
	admin, err := s.adminRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := util.GenerateJWT(admin.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *adminService) GetAdminByID(id int) (*database.Admin, error) {
	return s.adminRepo.FindByID(id)
}

func (s *adminService) UpdateAdmin(id int, req *admin.UpdateRequest) (*database.Admin, error) {
	admin, err := s.adminRepo.FindByID(id)
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

	err = s.adminRepo.Update(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *adminService) DeleteAdmin(id int) error {
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return err
	}
	return s.adminRepo.Delete(admin)
}

func (s *adminService) CreateUser(req *user.Request) error {
	user := &database.User{
		LineID:      "CreatedByAdmin",
		DisplayName: req.DisplayName,
		Phone:       req.Phone,
		IsMember:    req.IsMember,
	}
	return s.userRepo.Create(user)
}

func (s *adminService) GetUserByID(id int) (*database.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *adminService) UpdateUser(id int, req *user.Request) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	user.LineID = req.LineID
	user.DisplayName = req.DisplayName
	user.Email = req.Email
	user.LineToken = req.LineToken
	user.Phone = req.Phone
	user.IsMember = req.IsMember

	return s.userRepo.Update(user)
}

func (s *adminService) DeleteUser(id int) error {
	tx := s.userRepo.BeginTransaction()

	pendingOrders, err := s.orderRepo.FindPendingOrdersByUserIDTx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if len(pendingOrders) > 0 {
		tx.Rollback()
		return errors.New("user has pending orders, cannot delete")
	}

	err = s.userRepo.DeleteTx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
