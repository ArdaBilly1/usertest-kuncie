package service

import (
	"usertest-kuncie/src/gateway/request"
	"usertest-kuncie/src/model"
	"usertest-kuncie/src/repository"
)

type PromoService interface {
	CreatePromo(request *request.PromoInsertRequest) error
}

type promoService struct {
	repository repository.PromoRepository
}

func NewPromo(repo repository.PromoRepository) PromoService {
	return &promoService{repository: repo}
}

func (s *promoService) CreatePromo(request *request.PromoInsertRequest) error {
	var promo model.Promo

	promo.Name = request.Name
	promo.Type = request.Type
	promo.MinQty = request.MinQty
	promo.Percentage = request.Percentage
	promo.ItemBonusId = request.ItemBonusId

	return s.repository.Insert(promo)
}
