include .env

DOCKER=docker

APP_NAME=octo-tentacle

default: run

# Create a build docker image
build:
	$(DOCKER) build -t $(APP_NAME) .

# Remove built docker image
clean:
	- $(DOCKER) rm -f $(APP_NAME)

# After you build the image, check how it would run in production with this command
run: build
	$(DOCKER) run $(APP_NAME) $(NATS_SERVER)

