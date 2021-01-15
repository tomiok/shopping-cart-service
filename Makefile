# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=shopping-service

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	$(GOCMD) run main.go