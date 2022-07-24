package controllers

import (
	"github.com/kristabdi/bnmo-backend/models"
	"github.com/kristabdi/bnmo-backend/utils"
)

func UserInsertOne(data *models.User) error {
	result := utils.Db.Create(data)
	return result.Error
}

func UserGetByEmail(email string) (models.User, error) {
	var user = models.User{Email: email}

	result := utils.Db.First(&user)
	return user, result.Error
}

func UserGetBatch(page, pageSize int) ([]models.User, error) {
	var users []models.User

	result := utils.Db.Scopes(utils.Paginate(page, pageSize)).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := 0; i < len(users); i++ {
		users[i].NoPass()
	}

	return users, nil
}

func UserVerify(id uint) error {
	result := utils.Db.Model(&models.User{}).Where("id = ?", id).Update("is_verified", true)
	return result.Error
}

func UserUpdateBalance(id uint, balance uint64) error {
	result := utils.Db.Model(&models.User{}).Where("id = ?", id).Update("balance", balance)
	return result.Error
}
