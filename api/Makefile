.SILENT:

# find process on port (MacOS)
# netstat -vanp tcp | grep 8000

start:
	PX_ENV=development go run ./cmd/main/main.go

start\:db:
	docker-compose -f docker-compose.db.yml up -d --build

gql:
	go run github.com/99designs/gqlgen generate --config ./internal/http/graphql/gqlgen.yml
