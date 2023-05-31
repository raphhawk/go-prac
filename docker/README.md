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
```
