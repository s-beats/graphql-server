gen:
	gqlgen generate
dc-up:
	docker-compose up -d
dc-down:
	docker-compose down
start:
	go run server.go
exec-mysql:
	docker container exec -it mysql mysql -u root -p
