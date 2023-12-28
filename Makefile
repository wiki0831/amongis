
default:
	go build

build:
	docker compose build

run:
	Docker compose up -d

stop:
	Docker compose stop