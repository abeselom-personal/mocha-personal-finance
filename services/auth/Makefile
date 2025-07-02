MIGRATE ?= migrate
SWAGGER_OUT := ./docs/swagger

create:
	@timestamp=$$(date +%s); \
	name=$(name); \
	if [ -z "$$name" ]; then echo "Error: provide name='migration_name'"; exit 1; fi; \
	touch "./migrations/$${timestamp}_$$name.up.sql"; \
	touch "./migrations/$${timestamp}_$$name.down.sql"; \
	echo "Created migrations: $${timestamp}_$$name.up.sql and $${timestamp}_$$name.down.sql"

up:
	$(MIGRATE) -database $$DATABASE_URL -path ./migrations up

down:
	$(MIGRATE) -database $$DATABASE_URL -path ./migrations down

swagger-gen:
	swag init -g routes/routes.go -o $(SWAGGER_OUT)

.PHONY: create up down swagger-gen
