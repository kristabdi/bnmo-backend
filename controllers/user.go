package controllers

import (
	"github.com/kristabdi/bnmo-backend/models"
	"github.com/kristabdi/bnmo-backend/utils"
)

func UserInsertOne(data *models.User) error {
	result := utils.Db.Create(data)
	return result.Error
}

func UserGetByUsername(username string) (models.User, error) {
	user := models.User{Username: username}

	result := utils.Db.Where("username = ?", username).First(&user)
	return user, result.Error
}

func IDGetByUsername(username string) (uint, error) {
	user := models.User{Username: username}

	result := utils.Db.Where("username = ?", username).First(&user)
	return user.ID, result.Error
}

func UserGetAll() ([]models.User, error) {
	var users []models.User

	result := utils.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := 0; i < len(users); i++ {
		users[i].NoSensitive()
	}

	return users, nil
}

func UserVerify(username string) error {
	result := utils.Db.Model(&models.User{}).Where("username = ?", username).Update("is_verified", true)
	return result.Error
}

func UserUpdateBalance(id uint, balance uint64) error {
	result := utils.Db.Model(&models.User{}).Where("id = ?", id).Update("balance", balance)
	return result.Error
}
