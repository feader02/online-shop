include .env
export

migration_create:
	migrate create -ext sql -dir ./migrations -seq create_table

migration_up:
	migrate -path ./migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(MYSQL_TCP_PORT_EXPOSE))/$(DB_NAME)" up