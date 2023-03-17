include compose/.env
include common.mk

run: 
	DATABASE=$(database) \
	HOST=$(dbdomain).$(host) \
	PORT=$(dbport) \
	USER=$(dbuser) \
	PASSWORD=$(dbpassword) \
	SERVICEHOST=$(domain).$(host) \
	SERVICEPORT=$(port) \
	$(GORUN) cmd/main.go

.PHONY: run

