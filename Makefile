.PHONY: \
		build \
		start \
		stop \

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
