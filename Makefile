# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
BINARY_NAME=app
SOURCES=main.go index_id.go index_sex.go http.go


build:
		$(GOBUILD) $(SOURCES) -o $(BINARY_NAME) -v -a 

image:
		docker build -t stor.highloadcup.ru/accounts/clear_albatross .

push:
		docker push stor.highloadcup.ru/accounts/clear_albatross

run:
		$(GORUN) $(SOURCES)
