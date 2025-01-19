## Installing and Running VooSu Server Using Docker

## Requirements

Before starting, ensure you have the following tools installed:

- [Docker](https://docs.docker.com/get-started/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Installation Guide

Step 1: Clone the VooSu server repository with all its submodules

```bash
git clone --recursive https://github.com/voo-su/server.git
```

Step 2: Navigate to the 'server' directory

```bash
cd server
```

Step 3: Create the necessary directory structure for the web-client and clone the web-client repository

```bash
mkdir -p web/web-client && git clone https://github.com/voo-su/web.git web/web-client
```

Step 4: Set the correct permissions for the 'run-docker.sh' script

```bash
chmod 775 ./scripts/run-docker.sh
```

Step 5: Run the Docker setup script

```bash
./scripts/run-docker.sh
```
