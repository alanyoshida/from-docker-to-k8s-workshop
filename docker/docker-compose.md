# Docker compose

## How to Install

Install locally <https://docs.docker.com/compose/install/>

```yaml
version: "3.9"  # optional since v1.27.0
services:
  web:
    build: .
    ports:
      - "5000:5000"
    environment:
    - DEBUG=1
    volumes:
      - .:/code
      - logvolume01:/var/log
    links:
      - redis
  redis:
    image: redis
volumes:
  logvolume01: {}
```

### Commands

```bash
## Use local docker-compose.yaml to create containers
## -d for detach mode, run in background
docker-compose up -d

## Build the Dockerfiles
docker-compose build

## List the running containers
docker-compose ps

## Run a specific defined service in docker-compose.yaml
docker-compose run <service>

## Pause the containers
docker-compose stop

## Erase the containers
docker-compose down
```

### Reference

* [docker-compose](https://docs.docker.com/compose/compose-file/)
