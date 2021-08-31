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
start-redis:
	docker build -t redis .
	docker run --name redis -p 6379:6379 -d redis
stop-redis:
	docker stop redis
	docker rm redis