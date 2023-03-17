include compose/.env
include common.mk

run: 
	DATABASE=$(database) \
	HOST=$(localhost) \
	PORT=$(dbport) \
	USER=$(dbuser) \
	PASSWORD=$(dbpassword) \
	SERVICEHOST=$(localhost) \
	SERVICEPORT=$(port) \
	$(GORUN) cmd/main.go

docker-image:
	make -C docker docker-image

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(build) ./cmd/

docker-start:
	make -C compose start-db start-app

docker-stop:
	make -C compose stop

.PHONY: run

