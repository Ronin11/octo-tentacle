DOCKER=docker

APP_NAME=asdf

default: run

# Create a build docker image
build: clean
	$(DOCKER) build -t $(APP_NAME) .

# Remove built docker image
clean:
	- $(DOCKER) rm -f $(APP_NAME)

# After you build the image, check how it would run in production with this command
run: build
	$(DOCKER) run -p 80:80 $(APP_NAME)

