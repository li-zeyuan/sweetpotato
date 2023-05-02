EMPTY =
BUILD_TAG = $(shell git tag --contains=HEAD | tail -n 1)
ifeq ($(BUILD_TAG),$(EMPTY))
    BUILD_TAG = $(shell git show -s HEAD --format="%h")
endif

docker-build:
	docker build --rm --tag sweetpotato-highlightexam-backend:latest .
