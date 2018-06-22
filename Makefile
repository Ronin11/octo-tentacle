include .env

GO=go

APP_NAME=octo-tentacle

default: run

# Create a build go image
build:
	$(GO) build -o $(APP_NAME) .

# Remove built go image
clean:
	rm -f $(APP_NAME)

run: build
	export SERVER=$(NATS_SERVER); \
	./$(APP_NAME)

test:
	go test



