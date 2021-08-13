# Commands

```bash
## Pull image
docker pull <IMAGE>

## List Running Containers
docker ps

## List images
docker images

## List containers
docker container ls

docker stop <CONTAINER>

## Run a nginx container and attach the terminal to it
docker run -it <IMAGE> --name <CONTAINER_NAME>

## Eg:
docker run -it ubuntu
docker run -it ubuntu --name my-container

## Run with network host
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

## Execute command inside running container
docker exec -it <CONTAINER> <COMMAND>

```

## Volumes

<https://docs.docker.com/storage/volumes/>

```bash
docker run -v <HOST_FOLDER>:<CONTAINER_FOLDER>:ro <IMAGE>

## Eg:
docker run -v $(pwd):/tmp/my-foder:ro <IMAGE>

```

## Port forwarding

<https://docs.docker.com/config/containers/container-networking/>

```bash
# Run nginx with volume of current folder
docker run -v <HOST_FOLDER>:<CONTAINER_FOLDER>:ro -p <HOST_PORT>:<CONTAINER_PORT> <IMAGE>

## Eg:
cat << EOF > index.htlm
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>MINHA PAGINA HTML</title>
</head>
<body>
  <h1>TESTE DE MONTAGEM DE VOLUME EM NGINX</h1>
  <p>FUNCIONOU !!!</p>
</body>
</html>
EOF
docker run -v (pwd):/usr/share/nginx/html:ro -p 8181:80 nginx

## Show port forwarding from docker-proxy
sudo netstat -ntpl | grep 8181

## Test
curl -v http://localhost:8181

```
