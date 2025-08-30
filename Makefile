APP_NAME=go-book-management

# Generate Swagger docs
swagger:
	swag init -g main.go -o ./docs

# Build docker image
docker-build:
	docker build -t $(APP_NAME) .

# Run docker container
docker-run:
	docker run -p 8080:8080 --env-file .env -e DOCKER_ENV=true $(APP_NAME)

# Run everything: swagger + docker-build + docker-run
all: swagger docker-build docker-run