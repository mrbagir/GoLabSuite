# GoLabSuite

## 🚀 All-in-One Golang Project

Welcome to **GoLabSuite**, a monorepo project designed to collect, demonstrate, and combine various Golang technologies into a single, comprehensive project.

## 💡 Project Goals

- **Learning Center:** A knowledge hub and personal documentation for Golang technologies.
- **Reusable Template:** Ready-to-use boilerplates for different use cases.
- **Experiment Playground:** Open space to explore new technologies.

## ✅ Key Features
- Multi-service architecture: Online Shop, Rental Car, Exchange Rate, and more.
- Supports various protocols: HTTP, gRPC, GraphQL, MQTT.
- Cloud file management: AWS S3, Cloudinary, Minio.
- Payment Gateway Integration: Midtrans, Xendit, Omise.
- Database coverage: MySQL, PostgreSQL, MongoDB.
- Messaging and Event Streaming: RabbitMQ, Kafka, Redis.
- Deployment ready for Docker and Kubernetes.

## 📁 Project Structure
```
GoLabSuite/
├── cmd/                    # Main entry points for each service
├── internal/                # Business logic for each domain/service
├── api/                      # Protocol implementations (HTTP, gRPC, GraphQL, MQTT)
├── pkg/                      # Shared libraries (logger, middleware, utilities)
├── deployments/              # Deployment configurations (Docker, Kubernetes)
├── configs/                  # Configuration files (YAML, JSON)
├── docs/                      # Documentation files
├── tests/                     # Unit and integration tests
├── .env                       # Environment variables
├── Makefile                   # Task runner commands
├── go.mod                     # Go module definition
└── README.md                  # This file
```

## 📖 Usage Guide

### 1. Clone Repository
```bash
git clone https://github.com/mrbagir/GoLabSuite.git
```
### 2. Setup & Run
```bash
cp .env.example .env
make run
```
### 3. Run a Service (Example: Online Shop)
```bash
make run-onlineshop
```

## 📚 Detailed Documentation

Each module and mini-project has its own documentation within its respective folder.

## ✨ Contribution

If you have ideas to improve this project or want to add new features, feel free to open an issue or submit a pull request!
