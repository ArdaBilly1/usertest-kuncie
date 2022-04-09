package service

import "usertest-kuncie/src/repository"

type InventoryService interface {
}

type inventoryService struct {
	repository repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) InventoryService {
	return &inventoryService{repository: repo}
}

func (s *inventoryService) InsertInventory() {

}
