package storage

import (
	"fmt"

	"github.com/weiii576/tool/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewStore(env *configs.Env) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
