.SILENT:

# find process on port (MacOS)
# netstat -vanp tcp | grep 8000

start:
	PX_ENV=development go run ./services/project/cmd/main/main.go

start\:test:
	PX_ENV=test go run ./services/project/cmd/main/main.go

start\:db:
	docker-compose -f docker-compose.db.yml up -d --build

gql-core:
	go run github.com/99designs/gqlgen generate --config ./services/project/internal/delivery/http/graphql-core/gqlgen.yml

gql-client:
	go run github.com/99designs/gqlgen generate --config ./services/project/internal/delivery/http/graphql-client/gqlgen.yml
