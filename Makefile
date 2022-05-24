git_url := github.com/MatthiasScholz/temporal-tfsecret
init:
	go mod init $(git_url)

client:
	go build -o tfsecret-client ./cmd/start/main.go

worker:
	go build -o tfsecret-backend ./cmd/tfsecret-backend/main.go

deps:
	go mod tidy

version:
	make -version
	go version

clean:
	go clean -modcache

app:
	docker compose  --file docker-compose.app.yml up --detach

observability:
	docker compose --file docker-compose.observability.yml up --detach

temporal:
	docker compose --file docker-compose.temporal.yml up --detach

down-app:
	docker compose  --file docker-compose.app.yml down

down-observability:
	docker compose --file docker-compose.observability.yml down

down-temporal:
	docker compose --file docker-compose.temporal.yml down

up: temporal observability app

down: down-app down-observability down-temporal

restart: down up

ui:
	open http://localhost:8085
	open http://localhost:8088
