TAG?=latest

all: dockerapp
	docker build -t dajulia3/diego-edge-docker-app:${TAG} .
	rm dockerapp

dockerapp: main.go
	GOARCH=amd64 GOOS=linux go build -o dockerapp main.go
