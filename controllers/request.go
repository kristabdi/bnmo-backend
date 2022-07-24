package controllers

import (
	"github.com/kristabdi/bnmo-backend/models"
	"github.com/kristabdi/bnmo-backend/utils"
)

func RequestGetBatch(page int64, pageSize int64, id uint) ([]models.Request, error) {
	var requests []models.Request

	result := utils.Db.Scopes(utils.Paginate(int(page), int(pageSize))).Where("id = ?", id).Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}

	return requests, nil
}

func RequestInsertOne(data *models.Request) error {
	result := utils.Db.Create(data)
	return result.Error
}
