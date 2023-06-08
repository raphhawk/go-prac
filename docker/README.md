# Docker Cheatsheet

## Create Docker container (dedicated resource segments & ctags)

```sh
$ docker create <image name> <command>
```

## Start Docker container

```sh
$ docker start <container id>    # without logs
$ docker start -a <container id> # with logs
```

## Run Docker container

```sh
$ docker run <image name> <command>
```

## Get Docker containers created list

```sh
$ docker ps       # all running containers
$ docker ps --all # all containers created
```

## Dekete Docker containers

```sh
$ docker system prune # deletes entire containers
```

## Stop Docker containers

```sh
$ docker stop <container id>  # SIGTERM waiting for process
$ docker kill <container id>  # SIGKILL instant stop
```

## Execute additional commands on a running Docker container

```sh
# -it == -i -t, i means attach terminal to stdin of container, t means pretty format text from stdout

$ docker exec -it <container id>  <command> # run after starting the image

$ docker exec -it <container id>  sh # use terminal inside running container

$ docker run -it <image name> sh # use terminal directly after creating container
```

## Create a custom Docker image with existing image

### <li>create file named <i><b>Dockerfile</b></i> inside a new directory for the custom docker image</li>

```dockerfile
FROM alpine                 # specify image
RUN apk add --update redis  # download and install dependencies
CMD ["redis-server"]        # specify initialization command
```

### <li>execute following commands in the newly created directory</li>

```sh
$ docker build .
$ docker build -t raphhawk/redis:latest . # tagged image
```

### <li>creating images manually(not recommended)</li>

```sh
$ docker run -it alpine:latest sh
$ apk add --update redis                                 # run respective commands inside container
$ dockee commit -c 'CMD ["redis-server"]' <container id> # create image out of newly created container
```

## Run Docker images on ports(port forwarding)

```sh
$ docker run -p 8080:8080 <image name> # forward localhost:8080 => docker:8080
$ docker run -p 5000:8080 <image name> # forward localhost:5000 => docker:8080
```
