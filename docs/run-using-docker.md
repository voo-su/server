## Installing and Running VooSu Server Using Docker

## Requirements

Before starting, ensure you have the following tools installed:

- [Docker](https://docs.docker.com/engine/install/)

## Installation Guide

Step 1: Clone the VooSu server repository with all its submodules

```bash
git clone --recursive https://github.com/voo-su/server.git -b main
```

Step 2: Navigate to the 'server' directory

```bash
cd server
```

Step 3: Set the correct permissions for the 'run-docker.sh' script

```bash
chmod 775 ./scripts/run-docker.sh
```

Step 4: Run the Docker setup script

```bash
./scripts/run-docker.sh
```
