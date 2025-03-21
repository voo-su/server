#!/bin/bash

set -x

docker compose -f docker-compose.dev.yaml up -d

if ! docker inspect -f '{{.State.Running}}' postgres_dev &>/dev/null; then
    echo "The PostgreSQL container is not running."
    exit 1
fi

if ! docker inspect -f '{{.State.Running}}' clickhouse_dev &>/dev/null; then
    echo "The ClickHouse container is not running."
    exit 1
fi

until docker exec -i postgres_dev pg_isready -U postgres; do
    sleep 2
done

while IFS= read -r migration; do
    docker exec -i postgres_dev psql -U postgres -d voo_su <"$migration"
done < <(find ./migrations/postgres -type f -name "*.up.sql" | sort)

while IFS= read -r migration; do
    cat "$migration" | docker exec -i clickhouse_dev clickhouse-client --user clickhouse --password clickhouse --database=voo_su
done < <(find ./migrations/clickhouse -type f -name "*.up.sql" | sort)
