package routes

import (
	"gorm.io/gorm"
	"usertest-kuncie/src/gateway/handler"
	"usertest-kuncie/src/repository"
	"usertest-kuncie/src/service"
)

func ItemInjector(db *gorm.DB) handler.ItemHandler {
	repo := repository.NewItem(db)
	srv := service.NewItemService(repo)
	return handler.New(srv)
}
