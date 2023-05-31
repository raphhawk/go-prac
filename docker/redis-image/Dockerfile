# Use an existing docker image as a base
FROM alpine 
# Download and install a dependency
RUN apk add --update redis
# Specify what command to run on initial start of container
CMD ["redis-server"]

