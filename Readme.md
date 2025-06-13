# Fintech Application

## Project Structure
```
(include your directory structure here)
```

## Setup
1. Create `.env` file from template:
```bash
cp .env.example .env
```

2. Build and start services:
```bash
make setup
make deploy
```

## Commands

- `make build`  
  Build all docker services.

- `make up`  
  Start all services in detached mode.

- `make down`  
  Stop all running services.

- `make restart`  
  Restart services (down + up).

- `make logs`  
  Show recent logs and follow.

- `make clean`  
  Remove containers, images, and prune network.

- `make test`  
  Run tests inside running docker containers for each service.

- `make setup`  
  Initialize project (create network).

- `make deploy`  
  Build and start all services.

- `make status`  
  Show running containers status.

- `make auth-logs`  
  Follow logs of `auth` service.

- `make scraper-logs`  
  Follow logs of `receipt-scraper` service.

- `make build-services`  
  Build specific core services.

- `make interactive`  
  Interactive menu for running make commands.

- `make exec`  
  Open shell inside a running service container.

## Service Endpoints
- Auth: http://localhost:${GATEWAY_PORT_HTTP}/auth
- Notification: http://localhost:${GATEWAY_PORT_HTTP}/notification
- Receipt Scraper: http://localhost:${GATEWAY_PORT_HTTP}/receipt-scraper
- Regex: http://localhost:${GATEWAY_PORT_HTTP}/regex
- SMS Parser: http://localhost:${GATEWAY_PORT_HTTP}/sms-parser
- Sync: http://localhost:${GATEWAY_PORT_HTTP}/sync
```

### Key Improvements:
1. **Environment Management**:
   - Centralized `.env` file for all configurations
   - Automatic variable export in Makefile
   - Service-specific ports configurable in one place

2. **Network Configuration**:
   - Dedicated bridge network for secure inter-service communication
   - Automatic network creation via `make setup`

3. **Makefile Enhancements**:
   - Service-specific log targets
   - Clean system command
   - Setup initialization
   - Deployment pipeline
   - Status checking

4. **Gateway Configuration**:
   - Dynamic nginx config with environment variables
   - Proper service discovery via Docker network
   - SSL certificate support

5. **Operational Improvements**:
   - Containerized service isolation
   - Centralized logging
   - Health checks via status command
   - Rebuild automation

To use this setup:
1. Create `.env` file from the template
2. Run `make setup` to initialize network
3. Run `make deploy` to build and start all services

The system now supports:
- Service discovery via Docker network
- Dynamic port configuration
- Centralized logging
- One-command operations
- Scalable architecture
- Secure inter-service communication
