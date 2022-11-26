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
start: dc-up
	go run cmd/main.go

.PHONY: exec-mysql
exec-mysql: dc-up
	docker container exec -it graphql-todo-mysql mysql -u root -p

.PHONY: init-db
init-db: dc-up
	docker container exec graphql-todo-mysql sh /var/script/init_db.sh

.PHONY: generate-sqlc
generate-sqlc:
	sqlc generate
