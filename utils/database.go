package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type DbInstance struct {
	*gorm.DB
	*redis.Client
	context.Context
}

var Db DbInstance

func (Db *DbInstance) InitDB() error {
	var err error
	Db.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	Db.Context = context.Background()

	if err = Db.Client.Set(Db.Context, "IDR", 1.0, 20*time.Minute).Err(); err != nil {
		return err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DBNAME"), os.Getenv("DB_PORT"))

	Db.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err

}

func Paginate(page, pageSize int) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return Db.Offset(offset).Limit(pageSize)
	}
}
