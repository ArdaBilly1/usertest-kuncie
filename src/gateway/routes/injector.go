package routes

import (
	"usertest-kuncie/src/gateway/handler"
	"usertest-kuncie/src/repository"
	"usertest-kuncie/src/service"

	"gorm.io/gorm"
)

func ItemInjector(db *gorm.DB) handler.ItemHandler {
	repo := repository.NewItem(db)
	srv := service.NewItemService(repo)
	return handler.New(srv)
}

func PromoInjector(db *gorm.DB) handler.PromoHandler {
	repo := repository.NewPromo(db)
	srv := service.NewPromo(repo)
	return handler.NewPromoHandler(srv)
}

func OrderInjectory(db *gorm.DB) handler.OrderHandlerContract {
	repo := repository.NewOrder(db)
	repoDetail := repository.NewOrderDetail(db)
	srv := service.NewOrder(repo, repoDetail)
	return handler.NewOrderHandler(srv)
}
