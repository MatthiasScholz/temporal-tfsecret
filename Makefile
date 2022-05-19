git_url := github.com/MatthiasScholz/temporal-tfsecret
init:
	go mod init $(git_url)

worker:
	go build -v ./cli/tfsecret-backend -o tfsecret-backend

deps:
	go mod tidy

version:
	make -version
	go version
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

up: temporal observability

down: down-observability down-temporal

restart: down up

ui:
	open http://localhost:8085
	open http://localhost:8088
