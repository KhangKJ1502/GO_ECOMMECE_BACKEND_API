APP_NAME = server
CMD_PATH = ./cmd/$(APP_NAME)

run:
	go run $(CMD_PATH)

build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)
