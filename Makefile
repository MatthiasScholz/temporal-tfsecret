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
