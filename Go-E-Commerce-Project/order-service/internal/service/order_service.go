package service

import (
	"errors"

	"github.com/google/uuid"
	"order-service/internal/model"
	"order-service/internal/repository"
)

type OrderService interface {
	CreateOrder(userID uuid.UUID, items []OrderItemRequest) (*model.Order, error)
	GetOrderByID(id uuid.UUID) (*model.Order, error)
	GetOrdersByUser(userID uuid.UUID) ([]model.Order, error)
	UpdateOrderStatus(id uuid.UUID, status model.OrderStatus) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

type OrderItemRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

func (s *orderService) CreateOrder(userID uuid.UUID, items []OrderItemRequest) (*model.Order, error) {
	if len(items) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	order := &model.Order{
		ID:     uuid.New(),
		UserID: userID,
		Status: model.OrderPending,
		Items:  make([]model.OrderItem, len(items)),
	}

	var total float64
	for i, item := range items {
		// TODO: In real project → call Product Service (gRPC/HTTP) to get current price & check stock
		price := 99.99 // placeholder for now

		order.Items[i] = model.OrderItem{
			ID:        uuid.New(),
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     price,
		}
		total += price * float64(item.Quantity)
	}

	order.TotalAmount = total

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}

	// TODO: Publish OrderCreated event (RabbitMQ / Kafka) using shared later
	return order, nil
}

// Other methods...
func (s *orderService) GetOrderByID(id uuid.UUID) (*model.Order, error) {
	return s.repo.FindByID(id)
}

func (s *orderService) GetOrdersByUser(userID uuid.UUID) ([]model.Order, error) {
	return s.repo.FindByUserID(userID)
}

func (s *orderService) UpdateOrderStatus(id uuid.UUID, status model.OrderStatus) error {
	return s.repo.UpdateStatus(id, status)
}
