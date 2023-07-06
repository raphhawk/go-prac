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

# Kubernetes (k8s)

## Minikube

### use minikube for local k8s developement

```sh
$ minikube start                # start minikube
$ minikube ip                   # get minikube ip
$ eval $(minikube docker-env)   # execute docker commands inside minikube
```

## k8s commands

```sh
$ kubectl version                           # gets version of client & server
$ kubectl get pods                          # gets list of pods
$ kubectl get deployments                   # gets list of deployments
$ kubectl get services                      # gets list of setvices
$ kubectl apply -f <config.yaml>            # create pods/depls using yaml config
$ kubectl delete pod <pod name>             # delete pods
$ kubectl delete deployment <depl name>     # delete deployment
$ kubectl describe pod <pod name>           # describe pods
$ kubectl describe deployment <depl name>   # describe deployment
$ kubectl describe service <srv name>       # describe service
$ kubectl logs <pod name>                   # logs of particular pod
$ kubectl get namespace                     # gets list of available namespaces
$ kubectl get services -n <namespace name>  # get objects in a specific namespace 
```

## k8s communicating between objects of different namespaces
```url
http://<object_name>.<namespace_name>.src.cluster.local/...
```
