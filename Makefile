# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
BINARY_NAME=app
SOURCES=main.go index_id.go index_sex.go index_status.go http.go


build:
		$(GOBUILD) -o $(BINARY_NAME) $(SOURCES) 

image:
		docker build -t stor.highloadcup.ru/accounts/clear_albatross .

push:
		docker push stor.highloadcup.ru/accounts/clear_albatross

run:
		$(GORUN) $(SOURCES)

install:
		echo "Nothing to do here"

test:
		$(GOTEST) -v -cover

all: test
