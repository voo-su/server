#!/bin/bash

set -x

docker-compose up -d

if ! docker ps | grep -q clickhouse; then
    echo "The ClickHouse container is not running."
    exit 1
fi

if ! docker ps | grep -q postgres; then
    echo "The PostgreSQL container is not running."
    exit 1
fi

docker exec -i clickhouse clickhouse-client --user clickhouse --password clickhouse --database=voo_su < ./migrations/clickhouse/001_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/001_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/002_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/003_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/004_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/005_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/006_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/007_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/008_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/009_migration.up.sql
docker exec -i postgres psql -d voo_su -f ./migrations/postgresql/010_migration.up.sql
docker exec -it postgres psql -U postgres -d voo_su -c '\dt'
docker exec -it clickhouse clickhouse-client --user clickhouse --password clickhouse --database=voo_su -q 'SHOW TABLES;'
