set -eo pipefail

echo "Shopana: $APP_ENV"

mkdir -p "/root/apps/$APP_ENV/app"
mkdir -p "/root/apps/$APP_ENV/public"
mkdir -p "/root/apps/$APP_ENV/db-data" # TODO: create vanila db for prod

# cp -r ./.env "/apps/$APP_ENV/.env" TODO: inject variables
cp -r docker-compose.yml "/root/apps/$APP_ENV/app/docker-compose.yml"
cp -r nginx/nginx.conf "/root/apps/$APP_ENV/app/nginx.conf.template"

cd /root/apps/$APP_ENV/app

echo $DOCKER_LOGIN_PASSWORD | docker login ghcr.io -u reframework-bot --password-stdin

docker compose rm -f
docker compose --env-file .env up -d --build

exit 0
