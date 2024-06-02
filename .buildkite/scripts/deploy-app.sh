set -eou pipefail

echo "Shopana: $APP_ENV - $APP App"

cd "/apps/$APP_ENV/app"

echo $DOCKER_LOGIN_PASSWORD | docker login ghcr.io -u reframework-bot --password-stdin

docker compose pull $APP
docker compose up $APP -d --no-deps

exit 0
