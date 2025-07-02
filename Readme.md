# Fintech Application

## Project Overview
Containerized fintech application with microservices architecture. Includes authentication, SMS parsing, receipt scraping, and more - all managed through Docker and Make commands.

## Project Structure
```
fintech-app/
├── services/
│   ├── auth/             # Authentication service
│   ├── gateway/          # API Gateway (entry point)
│   ├── notification/     # Notification service
│   ├── receipt-scraper/  # Receipt processing
│   ├── regex/            # Regex pattern matching
│   ├── sms-parser/       # SMS parsing
│   └── sync/             # Data synchronization
├── Makefile              # Project commands
├── .env.example          # Environment template
└── docker-compose.yml    # Service orchestration
```

## 🚀 Quick Start

### 1. Initial Setup
```bash
# Copy environment file
cp .env.example .env

# Edit configuration (update ports/credentials)
nano .env

# Initialize project
make setup
```

### 2. Build & Run
```bash
# Build all services
make build

# Start all containers
make up

# Alternative: Single command build+start
make deploy
```

### 3. Verify Operation
```bash
# Check running services
make status

# View logs (all services)
make logs

# Test service functionality
make test
```

## ⚙️ Command Reference
| Command               | Description                                      |
|-----------------------|--------------------------------------------------|
| `make build`          | Build all Docker images                          |
| `make up`             | Start all services in background                 |
| `make down`           | Stop all running containers                      |
| `make restart`        | Restart all services (down + up)                 |
| `make logs`           | View combined logs (tail=100)                    |
| `make clean`          | Remove containers, images and networks           |
| `make test`           | Run tests with coverage reports                  |
| `make deploy`         | Full build + start pipeline                      |
| `make status`         | Show running container status                    |
| `make auth-logs`      | View authentication service logs                 |
| `make scraper-logs`   | View receipt scraper logs                        |
| `make build-services` | Rebuild core services only                       |
| `make interactive`    | Launch interactive command menu                  |
| `make exec`           | Access shell in running container                |

## 🔍 Service Endpoints
| Service           | Endpoint                                 |
|-------------------|------------------------------------------|
| Auth              | `http://localhost:${GATEWAY_PORT_HTTP}/auth` |
| Notification      | `http://localhost:${GATEWAY_PORT_HTTP}/notification` |
| Receipt Scraper   | `http://localhost:${GATEWAY_PORT_HTTP}/receipt-scraper` |
| Regex             | `http://localhost:${GATEWAY_PORT_HTTP}/regex` |
| SMS Parser        | `http://localhost:${GATEWAY_PORT_HTTP}/sms-parser` |
| Sync              | `http://localhost:${GATEWAY_PORT_HTTP}/sync` |
| **Swagger UI**    | `http://localhost:8081`                  |

## 🔄 Workflow Tips

### Testing Changes
```bash
# Rebuild specific services
make build-services

# Restart after code changes
make restart

# Run tests
make test
```

### API Documentation
```bash
# Generate Swagger docs
make swag-build

# Launch Swagger UI
make swagger-ui
```

### Debugging
```bash
# Follow specific service logs:
make auth-logs
make scraper-logs

# Access container shell:
make exec
# Then select service from list
```

## 🧹 Maintenance
```bash
# Full cleanup (containers, images, networks):
make clean

# Reinitialize project:
make setup
```

## Key Features
- **Single-command control** for all services
- **Automatic network configuration** with `make setup`
- **Service-specific logging** commands
- **Interactive command menu** for easy navigation
- **Automated Swagger docs** generation
- **Containerized testing** environment

> 💡 **Tip**: Use `make interactive` for a guided menu of all commands!
