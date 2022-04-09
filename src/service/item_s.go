package service

import (
	"database/sql"
	"fmt"
	"usertest-kuncie/src/gateway/request"
	"usertest-kuncie/src/gateway/response"
	"usertest-kuncie/src/model"
	"usertest-kuncie/src/repository"
)

type ItemService interface {
	CreateItem(request *request.InsertRequest) error
	GetAllItem() (items []response.ItemResponse, err error)
}

type itemService struct {
	repository repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{repository: repo}
}

var (
	ErrDataNotFound error = fmt.Errorf("data not found")
	ErrWhenCheckSku error = fmt.Errorf("sku must be unique")
)

func (s *itemService) CreateItem(request *request.InsertRequest) error {
	var item model.Item

	// check sku is unique
	idSku, err := s.repository.GetSKU(request.SKU)
	if err != sql.ErrNoRows || idSku > 0 {
		return ErrWhenCheckSku
	}

	item.ItemName = request.ItemName
	item.ItemPrice = request.ItemPrice
	item.SKU = request.SKU
	if err := s.repository.Create(item); err != nil {
		return err
	}

	return nil
}

func (s *itemService) GetAllItem() (items []response.ItemResponse, err error) {
	data, err := s.repository.GetAll()
	if err != nil {
		return items, err
	}

	if len(data) == 0 {
		return items, ErrDataNotFound
	}

	for i := range data {
		var item response.ItemResponse
		item.ID = data[i].ID
		item.ItemName = data[i].ItemName
		item.ItemPrice = data[i].ItemPrice
		item.SKU = data[i].SKU
		items = append(items, item)
	}

	return
}
