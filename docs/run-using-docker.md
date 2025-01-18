# Installing and Running VooSu Server Using Docker

## Requirements

Before starting, ensure you have the following tools installed:

- [Docker](https://docs.docker.com/get-started/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Installation Guide

Clone the repository with submodules:

```bash
git clone --recursive https://github.com/voo-su/server.git
```

Navigate to the project directory:

```bash
cd server
```

Create the required directory structure and clone the web-client repository:

```bash
mkdir -p web/web-client && git clone https://github.com/voo-su/web.git web/web-client
```

Start the server using Docker Compose:

```bash
docker-compose up -d
```
