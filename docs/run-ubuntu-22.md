## Installing and Running VooSu Server on Ubuntu 22.04

### Installing Additional Packages

``` sh
sudo apt update
sudo apt install make git curl
```

### PostgreSQL

``` sh
sh -c 'echo "deb https://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'

wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -

apt update && apt install postgresql-16 postgresql-client-16

systemctl status postgresql

psql --version

sudo -u postgres psql

ALTER USER postgres PASSWORD 'postgres';
```

### Redis

``` sh
add-apt-repository ppa:redislabs/redis

apt update && apt install redis

systemctl enable --now redis-server
```

### Clickhouse

``` sh
apt install apt-transport-https ca-certificates dirmngr

apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 8919F6BD2B48D754

echo "deb https://packages.clickhouse.com/deb stable main" | sudo tee /etc/apt/sources.list.d/clickhouse.list

apt update && apt install clickhouse-server clickhouse-client

service clickhouse-server start
```

### MinIO

``` sh
wget https://dl.min.io/server/minio/release/linux-amd64/minio_20250118003137.0.0_amd64.deb

dpkg -i minio_20250118003137.0.0_amd64.deb

groupadd -r minio-user

useradd -M -r -g minio-user minio-user

mkdir /mnt/data

chown minio-user:minio-user /mnt/data

nano /etc/default/minio
    MINIO_VOLUMES="/mnt/data"
    MINIO_OPTS="--address :9001 --console-address :9002"
    MINIO_ROOT_USER=
    MINIO_ROOT_PASSWORD=

systemctl start minio
systemctl status minio
systemctl enable minio
```

---

### NATS

``` sh
wget https://github.com/nats-io/nats-server/releases/download/v2.10.23-RC.4/nats-server-v2.10.23-RC.4-linux-amd64.tar.gz

tar -zxf nats-server-v2.10.23-RC.4-linux-amd64.tar.gz

cp nats-server-v2.10.23-RC.4-linux-amd64/nats-server /usr/bin/

nats-server -v

mkdir /etc/nats

nano /etc/nats/nats-server.conf
    cluster {
      name: "test-nats"
    }

    store_dir: "/var/lib/nats"
    listen: "0.0.0.0:4222"
    log_file: /var/log/nats/nats.log

useradd -r -c 'NATS service' nats
mkdir /var/log/nats /var/lib/nats
chown nats:nats /var/log/nats /var/lib/nats

nano /etc/systemd/system/nats-server.service
    [Unit]
    Description=NATS messaging server
    After=syslog.target network.target

    [Service]
    Type=simple
    ExecStart=/usr/bin/nats-server -c /etc/nats/nats-server.conf
    User=nats
    Group=nats
    LimitNOFILE=65536
    ExecReload=/bin/kill -HUP $MAINPID
    Restart=on-failure

    [Install]
    WantedBy=multi-user.target

systemctl enable nats-server --now
systemctl status nats-server
```

### Go

``` sh
https://go.dev/dl/go1.23.2.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz

# You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
export PATH=$PATH:/usr/local/go/bin
```

### Node.js

``` sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash

nvm install 22

corepack enable yarn

yarn -v
```

## Nginx

``` sh
apt install nginx
```

### Installation Guide

``` sh
git clone --recursive https://github.com/voo-su/server.git /usr/src/voo-su
```

``` sh
cd /usr/src/voo-su
```

``` sh
mkdir -p web/web-client && git clone https://github.com/voo-su/web.git web/web-client
```

``` sh
make install
```

``` sh
make build
```

``` sh
cd /usr/src/voo-su/web/web-client
```

``` sh
npm install --global yarn
```

``` sh
yarn install
```

You need to specify the domains in the file:

* /usr/src/voo-su/web/web-client/.env

``` sh
yarn build
```

``` sh
cd /usr/src/voo-su
```

``` sh
cp /usr/src/voo-su /usr/bin
```

``` sh
chmod 775 /usr/bin/voo-su
```

```bash
mkdit /etc/voo-su && ln -sf /usr/src/voo-su/configs/voo-su.yaml /etc/voo-su/voo-su.yaml
```

You need to configure the connection settings in the file:

* /usr/src/voo-su/configs/voo-su.yaml

```bash
voo-su cli-migrate
```

```bash
ln -sf /usr/src/voo-su/init/systemd/voo-su-http.service /etc/systemd/system/voo-su-http.service

ln -sf /usr/src/voo-su/init/systemd/voo-su-ws.service /etc/systemd/system/voo-su-ws.service

ln -sf /usr/src/voo-su/init/systemd/voo-su-cli-cron.service /etc/systemd/system/voo-su-cli-cron.service

ln -sf /usr/src/voo-su/init/systemd/voo-su-cli-queue.service /etc/systemd/system/voo-su-cli-queue.service

ln -sf /usr/src/voo-su/init/systemd/voo-su-grpc.service /etc/systemd/system/voo-su-grpc.service
```

```bash
systemctl daemon-reload
```

```bash
systemctl start voo-su-http.service voo-su-ws.service voo-su-cli-cron.service voo-su-cli-queue.service voo-su-grpc.service
```

```bash
systemctl enable voo-su-http.service voo-su-ws.service voo-su-cli-cron.service voo-su-cli-queue.service voo-su-grpc.service
```

You need to specify your subdomain in the Nginx configuration:

* /usr/src/voo-su/configs/nginx

```bash
ln -sf /usr/src/voo-su/configs/nginx /etc/nginx/sites-enabled/voo-su
```

```bash
systemctl restart nginx
```
