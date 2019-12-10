.PHONY: \
		build \
		start \
		stop \
		gen \
		test-start \
		test-stop \

SHELL := /bin/bash

# Builds the Docker container
build:
	docker build -t pmdcosta/coachee:latest .

# Starts an instance of the exchange
start:
	docker-compose up -d

# Stops the exchange instance
stop:
	docker-compose down

# Generates goa endpoints code
gen:
	goa gen coachee-backend/design

test-start:
	docker-compose --file docker-compose.local.yml up -d

test-stop:
	docker-compose --file docker-compose.local.yml down