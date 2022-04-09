package service

import (
	"fmt"
	"usertest-kuncie/src/gateway/request"
	"usertest-kuncie/src/model"
	"usertest-kuncie/src/repository"
)

type OrderService interface {
	Checkout(request *request.OrderRequest) error
}

type orderService struct {
	repository       repository.OrderRepositoryContract
	detailRepository repository.OrderDetailRepository
}

func NewOrder(
	repo repository.OrderRepositoryContract,
	detailRepo repository.OrderDetailRepository,
) OrderService {
	return &orderService{
		repository:       repo,
		detailRepository: detailRepo,
	}
}

var (
	ErrProductNotFound error = fmt.Errorf("product not found")
)

func (s *orderService) Checkout(request *request.OrderRequest) error {
	var order model.Order
	var idProducts []int

	for i := range request.OrderDetail {
		idProducts = append(
			idProducts,
			request.OrderDetail[i].ItemId,
		)
	}

	items, err := s.getProducts(idProducts)
	if err != nil {
		return ErrProductNotFound
	}

	order.CustomerName = request.CustomerName
	order.TotalAmount = s.calculateTotalAmount(items, request.OrderDetail)
	order.OrderStatus = model.OrderStatusInCart

	if err := s.repository.Insert(order); err != nil {
		return err
	}

	return nil
}

func (s *orderService) getProducts(ids []int) (item []model.Item, err error) {
	repo := s.repository.(*repository.OrderRepository)
	itemRepo := repository.NewItem(repo.DB)

	return itemRepo.GetByIds(ids)
}

func (s *orderService) calculateTotalAmount(
	items []model.Item,
	detailOrder []request.OrderDetailRequest,
) (totalAmount float32) {
	for i := range detailOrder {
		var price float32
		if items[i].ID == detailOrder[i].ItemId {
			price = float32(detailOrder[i].Qty) * items[i].ItemPrice
		}
		totalAmount += price
	}
	return
}
