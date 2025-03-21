#!/bin/bash

set -x

mkdir -p web/web-client && git clone https://github.com/voo-su/web.git web/web-client -d main

docker compose -f docker-compose.yaml up -d

if ! docker inspect -f '{{.State.Running}}' postgres &>/dev/null; then
    echo "The PostgreSQL container is not running."
    exit 1
fi

if ! docker inspect -f '{{.State.Running}}' clickhouse &>/dev/null; then
    echo "The ClickHouse container is not running."
    exit 1
fi

until docker exec -i postgres pg_isready -U postgres; do
    sleep 2
done

while IFS= read -r migration; do
    docker exec -i postgres psql -U postgres -d voo_su <"$migration"
done < <(find ./migrations/postgres -type f -name "*.up.sql" | sort)

while IFS= read -r migration; do
    cat "$migration" | docker exec -i clickhouse clickhouse-client --user clickhouse --password clickhouse --database=voo_su
done < <(find ./migrations/clickhouse -type f -name "*.up.sql" | sort)
