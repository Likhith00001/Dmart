# Dmart
E-Commerce platform


🛒 Go E-Commerce Microservices (Gin-Based)

A scalable, production-grade microservices architecture for an e-commerce platform built using Go (Golang) and Gin, following clean architecture, event-driven design, and cloud-native principles.

🚀 Architecture Overview

This project follows a Distributed Microservices Architecture with:

API Gateway (Gin)
Domain-based microservices
Event-driven communication (Kafka)
Database per service
Stateless services with JWT authentication
🔷 High-Level Flow
Client (Web / Mobile / Postman)
        ↓ HTTPS (JSON)
API Gateway (Gin)
        ↓
Microservices (Gin-based services)
        ↓
PostgreSQL (per service) + Redis (cache)
        ↓
Kafka (event bus)
        ↓
Background Workers
🧱 Core Design Principles
Clean Architecture
Handler → Service → Repository
Loose Coupling via Events (Kafka)
Database per Service
Eventual Consistency using Saga Pattern
Stateless Services with JWT Authentication
Resilience
Retries
Circuit Breakers
Idempotency
Observability
Structured Logging (Zap)
Correlation IDs
Distributed Tracing
🏗️ Services
Service	Responsibility
API Gateway	Routing, Auth, Middleware
User Service	Authentication & Profile
Product Service	Product Catalog
Cart Service	Shopping Cart
Order Service	Order Management
Notification Service	Event-based Notifications
(Future) Payment Service	Payment Processing
(Future) Inventory Service	Stock Management
⚙️ Tech Stack
Layer	Technology
Language	Go 1.23+
HTTP Framework	Gin
Database	PostgreSQL
Cache	Redis
Messaging	Kafka
Auth	JWT
Logging	Zap
Config	Viper + .env
Containerization	Docker
Orchestration	Kubernetes
Migrations	golang-migrate
Validation	validator.v10
Observability	OpenTelemetry
📁 Project Structure
🗂 Monorepo Layout
go-ecommerce-microservices/
├── api-gateway/
├── user-service/
├── product-service/
├── cart-service/
├── order-service/
├── notification-service/
├── docker-compose.yml
├── k8s/
├── shared/
└── README.md
🔧 Per-Service Structure
user-service/
├── cmd/
│   └── server/main.go
├── internal/
│   ├── config/
│   ├── middleware/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   ├── cache/
│   ├── queue/
│   ├── worker/
│   └── utils/
├── pkg/
│   └── logger/
├── migrations/
├── docker/
├── test/
├── go.mod
├── .env.example
└── README.md
🌐 API Gateway Responsibilities
Request Routing
JWT Authentication
Rate Limiting (Redis)
Logging & Monitoring
CORS Handling
Panic Recovery
🔗 Route Structure
/api/v1/auth/*
/api/v1/products/*
/api/v1/cart/*
/api/v1/orders/*
🔐 Authentication Flow
User logs in via /auth/login
JWT is issued
Client sends token in headers
API Gateway:
Validates JWT
Injects user info into context
Services consume user context
🔄 Communication Patterns
✅ Synchronous
HTTP/REST (Gin)
Optional gRPC (internal high-performance calls)
✅ Asynchronous
Kafka Events
UserCreated
ProductUpdated
OrderCreated
🧠 Caching Strategy
Cache-Aside Pattern
Redis used for:
Product data
Cart state
Rate limiting
🔁 Event-Driven Workflow (Example)

Order Placement Flow:

User places order
Order Service:
Creates order
Publishes OrderCreated
Other services react:
Notification Service → Sends alert
Inventory Service → Updates stock
Payment Service → Processes payment
🛠️ Implementation Roadmap
📍 Phase 0: Foundation
Setup Go modules
Viper config + .env
Zap logging
PostgreSQL setup
Basic Gin server
Docker setup
📍 Phase 1: API Gateway
Middleware chain:
Request ID
Logging
Recovery
CORS
Rate Limiting
JWT Auth
📍 Phase 2: User Service
Register / Login
JWT generation
Protected routes
📍 Phase 3: Product Service
CRUD APIs
Filtering & Pagination
Kafka event publishing
📍 Phase 4: Cart Service
Add/Update/Delete cart items
Transaction handling
📍 Phase 5: Order Service
Order creation
Saga orchestration
Idempotency
📍 Phase 6: Redis Caching
Cache-aside implementation
📍 Phase 7–8: Kafka & Workers
Consumers
Retry logic
Background jobs
📍 Phase 9: Production Readiness
Error handling standardization
OpenAPI docs
Circuit breakers
Observability
📍 Phase 10: Testing & Deployment
Unit & integration tests
Docker optimization
Kubernetes manifests
🧪 Testing Strategy
Unit Tests (Service Layer)
Integration Tests (DB + API)
API Testing (Postman / Curl)
Load Testing (future)
🐳 Running the Project
🔹 Using Docker Compose
docker-compose up --build

Services:

API Gateway → :8080
User Service → :8081
Product Service → :8082
Cart Service → :8083
Order Service → :8084
🩺 Health Check
GET /health

Response:

{
  "status": "ok"
}
📊 Observability
Structured Logs (Zap)
Correlation IDs across services
Distributed tracing (OpenTelemetry)
💡 Gin Best Practices Used
gin.New() for full control
Custom middleware chain
ShouldBindJSON() for validation
Context-based user injection
Centralized error handling
📌 Future Enhancements
gRPC internal communication
Service discovery (Consul / etcd)
GraphQL Gateway
Advanced rate limiting (AI-based)
Multi-region deployment
🤝 Contribution

This project is designed for:

Learning microservices in Go
Interview preparation
Real-world backend system design
📄 License

MIT License
