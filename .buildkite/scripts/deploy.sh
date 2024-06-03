set -eo pipefail

echo "Shopana: $APP_ENV"

DOCKER_PASSWORD=$(buildkite-agent secret get docker_login_password)

echo "Checking out"
cd repo

echo "Building"

mkdir -p "~/apps/$APP_ENV/app"
mkdir -p "~/apps/$APP_ENV/public"
mkdir -p "~/apps/$APP_ENV/db-data" # TODO: create vanila db for prod

bash .buildkite/scripts/generate-env.sh

cp -r .env                 "~/apps/$APP_ENV/app/.env"
cp -r docker-compose.yml   "~/apps/$APP_ENV/app/docker-compose.yml"
cp -r nginx/nginx.conf     "~/apps/$APP_ENV/app/nginx.conf.template"

cd "~/apps/$APP_ENV/app"

echo $DOCKER_PASSWORD | docker login ghcr.io -u reframework-bot --password-stdin

docker compose rm -f
docker compose --env-file .env up -d --build

exit 0
