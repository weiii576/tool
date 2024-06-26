package user

import (
	"math"

	"github.com/weiii576/tool/internal"
	"gorm.io/gorm"
)

type (
	UserStorage interface {
		CreateUser(*User) error
		GetUsers(*internal.PaginationRequest) (*[]Profile, *internal.PaginationResponse, error)
		GetUserByEmail(string) (*User, error)
	}
	userStorage struct {
		Store *gorm.DB
	}
)

func NewUserStorage(store *gorm.DB) UserStorage {
	return &userStorage{
		Store: store,
	}
}

func (uc *userStorage) CreateUser(user *User) error {
	if result := uc.Store.Create(user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (uc *userStorage) GetUsers(pagination *internal.PaginationRequest) (*[]Profile, *internal.PaginationResponse, error) {
	var users = &[]User{}
	var profiles = &[]Profile{}
	var count int64

	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if pagination.PerPage == 0 {
		pagination.PerPage = 10
	}

	// get all users
	query := uc.Store.Model(users)

	// search for fields
	if pagination.Search != "" {
		searchPattern := "%" + pagination.Search + "%"
		query = query.Where("name LIKE ? OR email LIKE ?", searchPattern, searchPattern)
	}

	err := query.Count(&count).Error
	if err != nil {
		return nil, nil, err
	}

	maxPage := int64(math.Ceil(float64(count) / float64(pagination.PerPage)))

	// select fields exist in Profile
	result := query.Limit(pagination.PerPage).Offset(pagination.GetOffset()).Find(profiles)

	return profiles, &internal.PaginationResponse{
		Page:    pagination.Page,
		PerPage: pagination.PerPage,
		MaxPage: maxPage,
		Count:   count,
	}, result.Error
}

func (uc *userStorage) GetUserByEmail(email string) (*User, error) {
	var user = &User{}
	result := uc.Store.Where("email = ?", email).First(&user)

	return user, result.Error
}
