package user

import (
	"github.com/weiii576/tool/storage"
)

type UserStorage struct {
	Store *storage.PostgresStore
}

func NewUserStorage(store *storage.PostgresStore) *UserStorage {
	return &UserStorage{
		Store: store,
	}
}

func (uc *UserStorage) CreateUser(user *User) error {
	if result := uc.Store.DB.Create(user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (uc *UserStorage) GetUserByEmail(email string) (*User, error) {
	var user = &User{}
	result := uc.Store.DB.Where("email = ?", email).First(&user)

	return user, result.Error
}
