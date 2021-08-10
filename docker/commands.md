# Commands

```bash
## List Running Containers
docker ps

## List images
docker images

## Run a nginx container and attach the terminal to it
docker run -it ubuntu
docker run -it --net host alpine

## Build the Dockerfile in current folder
docker build .

## Build passing the Dockerfile path
docker build -f /path/to/a/Dockerfile .

## Build and tag the image
docker build -t shykes/myapp .

## Inspect the Image
docker image inspect

## Show container configurations
docker inspect <CONTAINER>

## Show logs from a container
docker logs <CONTAINER>

## Show mem/cpu usage
docker stats <CONTAINER>
```
