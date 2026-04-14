package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderPending   OrderStatus = "PENDING"
	OrderConfirmed OrderStatus = "CONFIRMED"
	OrderShipped   OrderStatus = "SHIPPED"
	OrderDelivered OrderStatus = "DELIVERED"
	OrderCancelled OrderStatus = "CANCELLED"
)

type Order struct {
	ID          uuid.UUID   `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID      uuid.UUID   `gorm:"type:uuid;not null;index"`
	TotalAmount float64     `gorm:"not null"`
	Status      OrderStatus `gorm:"type:varchar(20);not null;default:'PENDING'"`
	Items       []OrderItem `gorm:"foreignKey:OrderID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"not null"` // price at the time of ordering
	CreatedAt time.Time
}
