package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "order-service/internal/model"
    "order-service/internal/service"
)

type OrderHandler struct {
    service service.OrderService
}

func NewOrderHandler(svc service.OrderService) *OrderHandler {
    return &OrderHandler{service: svc}
}

type CreateOrderRequest struct {
    Items []service.OrderItemRequest `json:"items" binding:"required,min=1,dive"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
    userIDStr := c.GetString("user_id") // from JWT middleware
    userID, _ := uuid.Parse(userIDStr)

    var req CreateOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order, err := h.service.CreateOrder(userID, req.Items)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
        return
    }

    order, err := h.service.GetOrderByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}

// Add GetUserOrders, UpdateStatus etc. similarly