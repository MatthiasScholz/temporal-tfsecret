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

directory := internal/pkg/activities
test: deps
	 go test ./...

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

up: temporal observability ps

down: down-observability down-temporal ps

restart: down up ps

ps:
	docker compose --file docker-compose.observability.yml ps
	docker compose --file docker-compose.temporal.yml ps

logs:
	docker compose --file docker-compose.observability.yml logs
	docker compose --file docker-compose.temporal.yml logs

ui:
	open http://localhost:8085
	open http://localhost:8088
