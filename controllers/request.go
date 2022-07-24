package controllers

import (
	"github.com/google/uuid"
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

func RequestVerify(id uuid.UUID) error {
	result := utils.Db.Model(&models.Request{}).Where("id = ?", id).Update("is_approved", true)
	return result.Error
}

func RequestGetAll() ([]models.Request, error) {
	var req []models.Request

	result := utils.Db.Find(&req)
	if result.Error != nil {
		return nil, result.Error
	}

	return req, nil
}

func RequestGetById(id uuid.UUID) (models.Request, error) {
	req := models.Request{ID: id}

	result := utils.Db.Where("id = ?", id).First(&req)
	return req, result.Error
}
