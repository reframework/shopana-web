#!/bin/bash

create_docker_secret() {
    local secret_name=$1
    local secret_value=$2
    echo "$secret_value" | docker secret create --label external=true "$secret_name" -
}

get_buildkite_secret() {
    local secret_name=$1
    local secret_value=$(buildkite-agent secret get "$secret_name")
    echo "$secret_value"
}

if [ "$#" -lt 1 ]; then
    echo "Usage: $0 <secret_name1> <secret_name2> ... <secret_nameN>"
    exit 1
fi

for secret in "$@"; do
    # Get the value from Buildkite agent
    secret_value=$(get_buildkite_secret "$secret")

    if [ -z "$secret_value" ]; then
        echo "Buildkite secret $secret is not set. Skipping..."
        exit 1
        continue
    fi

    create_docker_secret "$secret" "$secret_value"
done

echo "==============================="
echo "Docker secrets creation process completed."
echo "==============================="
