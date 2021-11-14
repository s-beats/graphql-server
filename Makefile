gen:
	gqlgen generate
dc-up:
	docker-compose up -d
dc-down:
	docker-compose down
start:dc-up
	go run cmd/main.go
exec-mysql:
	docker container exec -it mysql mysql -u root -p
