# Makefile
.PHONY: build up down restart logs clean test setup

ENV_FILE := .env
GREEN=\033[0;32m
RED=\033[0;31m
YELLOW=\033[0;33m
NC=\033[0m
include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

restart:
	@make down
	@make up

logs:
	docker-compose logs -f --tail=100

clean:
	docker-compose down -v --rmi all
	docker network prune -f


test:
	@echo "Running all service tests with coverage inside docker containers..."
	@total=0; passed=0; skipped=0; \
	for dir in $$(find services -name go.mod -exec dirname {} \;); do \
		service=$$(basename $$dir); \
		echo "Testing $$dir inside container $$service..."; \
		cont=$$(docker-compose ps -q $$service); \
		if [ -z "$$cont" ]; then \
			printf "\033[0;33m⚠️ Container for %s not running, skipping tests\033[0m\n" "$$service"; \
			skipped=$$((skipped+1)); \
			continue; \
		fi; \
		pkgs=$$(docker-compose exec -T $$service sh -c "cd /app && go list ./..." 2>/dev/null); \
		if [ -z "$$pkgs" ]; then \
			printf "\033[0;32m✅ No tests found in %s\033[0m\n" "$$dir"; \
			skipped=$$((skipped+1)); \
			continue; \
		fi; \
		out=$$(docker-compose exec -T $$service sh -c "cd /app && go test -cover -short -count=1 ./..." 2>&1); \
		echo "$$out" | grep -E 'PASS|FAIL' || true; \
		if echo "$$out" | grep -q FAIL; then \
			printf "\033[0;31m❌ %s failed\033[0m\n" "$$dir"; \
		else \
			printf "\033[0;32m✅ %s passed\033[0m\n" "$$dir"; \
			passed=$$((passed+1)); \
		fi; \
		total=$$((total+1)); \
	done; \
	echo "---------"; \
	printf "\033[0;33mPassed: %d / %d, Skipped: %d\033[0m\n" $$passed $$total $$skipped; \
	if [ $$passed -ne $$total ]; then exit 1; fi

setup: 
	@echo "Initializing project..."
	docker network create ${NETWORK_NAME} || true
	@echo "Project setup complete"

deploy: build up

status:
	docker-compose ps

# Service-specific targets
auth-logs:
	docker-compose logs -f auth

scraper-logs:
	docker-compose logs -f receipt-scraper

build-services:
	docker-compose build \
		auth \
		notification \
		receipt-scraper \
		regex \
		sms-parser \
		sync \
		gateway

interactive:
	@echo "Choose a target:"
	@select opt in build up down restart logs clean test setup deploy status auth-logs scraper-logs build-services; do \
		if [ -n "$$opt" ]; then \
			echo "Running: make $$opt"; \
			make $$opt; \
			break; \
		else \
			echo "Invalid option"; \
		fi \
	done

exec:
	@echo "Select a service to exec into:"
	@docker-compose ps --services | nl -w2 -s'. '
	@read -p "Enter number: " choice; \
	 service=$$(docker-compose ps --services | sed -n "$${choice}p"); \
	 echo "Opening shell in $$service..."; \
	 docker-compose exec $$service sh || docker-compose exec $$service bash
