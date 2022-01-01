.PHONY: gqlgen
gqlgen:
	gqlgen generate

.PHONY: dc-up
dc-up:
	docker-compose up -d

.PHONY: dc-down
dc-down:
	docker-compose down

.PHONY: start
start:dc-up
	go run cmd/main.go

.PHONY: exec-mysql
exec-mysql:
	docker container exec -it mysql mysql -u root -p

.PHONY: init-db
init-db:
	docker container exec mysql sh /var/script/init_db.sh