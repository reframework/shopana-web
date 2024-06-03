#!/bin/bash

echo "Shopana: $APP_ENV"

DOCKER_PASSWORD=$(buildkite-agent secret get docker_login_password)

echo "Building"

mkdir -p "$WORKDIR/app"
mkdir -p "$WORKDIR/public"
mkdir -p "$WORKDIR/db-data"

bash .buildkite/scripts/generate-env.sh

# List of secret names
secret_names=(
  "cms_guset_jwt"
  "cms_customer_jwt"
  "cms_tenant_jwt"
  "cms_db_dns"
  "webapi_jwt_secret"
  "webapi_db_dns"
  "webclient_rollbar_server_token"
)

bash .buildkite/scripts/generate-secrets.sh "${secret_names[@]}"

cp -r .env                 "$WORKDIR/app/.env"
cp -r docker-compose.yml   "$WORKDIR/app/docker-compose.yml"
cp -r nginx/nginx.conf     "$WORKDIR/app/nginx.conf.template"

cd "$WORKDIR/app"

echo $DOCKER_PASSWORD | docker login ghcr.io -u reframework-bot --password-stdin

docker compose rm -f
docker compose --env-file .env up -d --build

exit 0
