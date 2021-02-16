migrateup:
	migrate -path db/migrations -database "${DATABASE_CONNECTION_STRING}" -verbose up

migratedown:
	migrate -path db/migrations -database "${DATABASE_CONNECTION_STRING}" -verbose down

.PHONY: migrateup migratedown