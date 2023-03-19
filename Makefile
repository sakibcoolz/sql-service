include compose/.env
include variables.mk

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
	make -C compose start

docker-stop:
	make -C compose stop

.PHONY: run

