package controllers

import (
	"github.com/kristabdi/bnmo-backend/models"
	"github.com/kristabdi/bnmo-backend/utils"
)

func TransactionGetBatch(page int64, pageSize int64, id uint) ([]models.Transaction, error) {
	var transactions []models.Transaction

	result := utils.Db.Scopes(utils.Paginate(int(page), int(pageSize))).Where("transactions.id_from = ?", id).Or("transactions.id_to = ?", id).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := 0; i < len(transactions); i++ {
		transactions[i].NameFrom = transactions[i].UserFrom.Name
		transactions[i].NameTo = transactions[i].UserTo.Name
	}

	return transactions, nil
}

func TransactionInsertOne(data *models.Transaction) error {
	result := utils.Db.Create(data)
	return result.Error
}
