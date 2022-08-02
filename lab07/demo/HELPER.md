# Extra Information

## Introduction

So far we've seen how we write our first go apps in a local environment and how we
interact with it. Today we will see how we can do this in a manner that allows us to easily replicate the setup on other hosts without any difficulty.

## Setup

### Toolstack

For this lab you will be required to have docker installed, therefore you will need to
follow the guide mentioned [here](https://docs.docker.com/get-docker/) in order to install it. This should also install the `docker-compose` plugin that we are going to use later.

An easier alternative is to create a docker account and use [docker playground](https://labs.play-with-docker.com/) which comes along with all the toolchain required for this lab, only be mindful that your session is limited to 4 hours.

### Hello world

After configuring your workspace one way to validate your setup is to run the `hello-world` image:
```
$ docker run hello-world:latest
```

If everything is configured properly the output should be something like this:
```
Hello from Docker!
This message shows that your installation appears to be working correctly.
```

## Docker CLI

### Running a shell inside a container

For the next exercise we are going to start shell inside an Ubuntu:22.04 container as follows:

```
$ docker run -it --name shell ubuntu:22.04 bash
root@bd60b3f8e067:/# pwd
/
root@bd60b3f8e067:/# whoami
root
```

The presented parameters are doing the following:

- `run` tells docker to start a container
- `-i` starts the container in interactive mode which instructs the container to accept input from STDIN
- `-t` associates a terminal with your container
- `--name` provides a name for you container
- `ubuntu:22.04` (`<image-name>:<tag>`) represents the image used to launch my container
- `bash` represents the command executed by the container, or the `entrypoint`

You can exit using `Ctrl+D` or type `exit` in the terminal.

> **_NOTE_** You should keep in mind that a container's life span is determined by its entrypoint execution time. This means that if we provide our container with a command like `echo "Hey I know Docker now"` it will execute it and then just exit and enter in a stopped state.

That being said, if we try to print the previous containers with `docker ps` the output will look like this:
```
$ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```

Adding the parameter `-a` will instruct docker to list all containers including the stopped ones:
```
$ docker ps -a
CONTAINER ID   IMAGE          COMMAND      CREATED             STATUS                            PORTS     NAMES
bd60b3f8e067   ubuntu:22.04   "bash"       44 minutes ago      Exited (0) 42 minutes ago                   shell
76c9dcacd3f1   hello-world    "/hello"     About an hour ago   Exited (0) About an hour ago                thirsty_wing
```

Use `docker rm <container-id-list-space-separated>` to clean your host (e.g. `docker rm bd60b3f8e067 76c9dcacd3f1`).

If we want to have that `bash` always available we can try something a little bit different that mimics a `while(true)` expression as an entrypoint and also we would be required to start the container in `detached mode` using `-d`, otherwise we would only be able to see the output from the entrypoint that would run indefinitely. For this example we are going to use `sleep infinity` that will wait indefinitely until we issue `Ctrl+C`:
```
$ docker run -d ubuntu:22.04 sleep infinity
```

Afterwards we can see that the container is still up and running:
```
$ docker ps -a
CONTAINER ID   IMAGE          COMMAND            CREATED         STATUS         PORTS     NAMES
4dd84d198721   ubuntu:22.04   "sleep infinity"   4 seconds ago   Up 3 seconds             pensive_solomon
```

And in order to get a shell we can use the `exec` parameter that will instruct docker to execute a given command on a running container:
```
$ docker exec -it pensive_solomon bash
root@4dd84d198721:/#
root@4dd84d198721:/#
```

This could prove to be quite useful when it comes to debugging a running service(e.g. you can imagine that we have a service which is unable to connect to another component and in order to debug that connectivity issue we can connect to the container and use various tools to see where the problem is).

> **_NOTE_** If you want to clean fast your entire environment and remove all the containers you can use the following command: `docker stop $(docker ps -aq) && docker rm $(docker ps -aq)`

## Docker Images

Now that you've seen how we are working with `run` and `exec` lets see how we can create a custom image that suits our needs to deploy an app. For this section we will work with dockerfiles.

### What is a Dockerfile?

A Dockerfile is basically a set of commands that are instructing docker on how to build and image (e.g. you can instruct it to copy some configuration files for an apache server).

### Example app

To avoid losing focus on Docker’s features, the sample application is a minimal HTTP server that has only three features:

- It responds with a text message containing a heart symbol (“<3”) on requests to /.
- It responds with {"Status" : "OK"} JSON to the health check request on requests to /ping.
- The port it listens on is configurable using the environment variable HTTP_PORT. The default value is 8080.
```
package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
```
### Creating a Docker Image

An image is defined by a file called `Dockerfile`, which specifies what happens inside the container we want to create, where access to resources (such as network interfaces or hard drives) is virtualized and isolated from the rest of the system. Through this file, we can specify port mappings, files that will be copied to the container when it starts, etc.

A Dockerfile is similar to a Makefile, and each line in it describes a layer in the image. Once we've defined a correct Dockerfile, our application will always behave **identically** no matter what environment it's run on. An example Dockerfile for our application is as follows:

```
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
```
- **FROM**: tells Docker what base image we would like to use for our application.
- **WORKDIR**: creates a directory inside the image that we are building. Instructs Docker to use this directory as the default destination for all subsequent commands.
- **COPY**: takes two parameters: the first parameter tells Docker what files you want to copy into the image and the last parameter tells Docker where you want that file to be copied to.
> **_NOTE_** Before we can run go mod download inside our image, we need to get our go.mod and go.sum files copied into it.
- **RUN**: execute a given command.
> **_NOTE_** "RUN go mod download" works exactly the same as if we were running locally on our machine, but this time these Go modules will be installed into a directory inside the image.
- **EXPOSE**: exposes a port outside the container.
- **CMD**: tell Docker what command to execute when our image is used to start a container.
> **_NOTE_** The EXPOSE statement does not actually expose the port given as a parameter, but functions as a kind of documentation between whoever is building the image and whoever is running the container, about which ports should be published. To publish a port when running a container, use the -p flag on the docker run command.

### Building an image
Now that we’ve created our Dockerfile, let’s build an image from it. The docker build command creates Docker images from the Dockerfile and a “context”. A build context is the set of files located in the specified path or URL. The Docker build process can access any of the files located in the context.

The build command optionally takes a `--tag` flag. This flag is used to label the image with a string value, which is easy for humans to read and recognise. If you do not pass a --tag, Docker will use latest as the default value.
```
 $ docker build --tag docker-gs-ping .
     [internal] load build definition from Dockerfile                                      0.1s
     => transferring dockerfile: 38B                                                       0.0s
     [internal] load .dockerignore                                                         0.1s
     => transferring context: 2B                                                           0.0s
     [internal] load metadata for docker.io/library/golang:1.16-alpine                     3.0s
     [1/7] FROM docker.io/library/golang:1.16-alpine@sha256:49c07aa83790aca732250c2258b59  0.0s
     => resolve docker.io/library/golang:1.16-alpine@sha256:49c07aa83790aca732250c2258b59  0.0s
     [internal] load build context                                                         0.1s
     => transferring context: 114B                                                         0.0s
     CACHED [2/7] WORKDIR /app                                                             0.0s
     CACHED [3/7] COPY go.mod .                                                            0.0s
     CACHED [4/7] COPY go.sum .                                                            0.0s
     CACHED [5/7] RUN go mod download                                                      0.0s
     CACHED [6/7] COPY *.go .                                                              0.0s
     CACHED [7/7] RUN go build -o /docker-gs-ping                                          0.0s
     exporting to image                                                                    0.1s
     => exporting layers                                                                   0.0s
     => writing image sha256:336a3f164d0f079f2e42cd1d38f24ab9110d47d481f1db7f2a0b0d2859ec  0.0s
     => naming to docker.io/library/docker-gs-ping                                         0.0s
```

### Useful commands
 - List local images:
   - `docker image ls`
   - `docker images`
 - List details about an image: 
   - `docker image inspect appname`
 - Run an image
   - `docker container run -p exposed_port:internal_port appname`
      ```
     $ docker container run -p 8888:8080 docker-gs-ping

     * Running on http://0.0.0.0:8080/ (Press CTRL+C to quit)
      172.17.0.1 - - [23/Sep/2019 14:46:00] "GET / HTTP/1.1" 200 -
      172.17.0.1 - - [23/Sep/2019 14:46:01] "GET /favicon.ico HTTP/1.1" 404 -
      172.17.0.1 - - [23/Sep/2019 14:46:02] "GET / HTTP/1.1" 200 –
      [...]
     ```
By accessing http://127.0.0.1:8888 from a web browser, we will see the web application we created.

The `-p` flag exposes the application's port 8080 and specifies a mapping between it and port 8888 on the machine we're running on. If we want to run the application in detached mode, we can do it using the `-d` flag.

## Mounts

Now that we know how to write a docker image and deploy it, we need to start addressing the persistence problem, in order to persist our data between different runs.

We are going to address this problem using mounts, specifically: `bind mounts` and `docker volumes`.

### Bind mounts

Basically those mounts are allowing a host to share its filesystem between multiple containers with read-only or read-write permission. And are quite useful when we want to configure one of our services when we deploy a container. As an example lets say we have the following postgresql script that will create our database with an user called `student`:

```
CREATE USER student;
CREATE DATABASE workshop;
GRANT ALL PRIVILEGES ON DATABASE workshop TO student;
```

One way to do that is to mount a script with suffix `.sql` in `/docker-entrypoint-initdb.d/` using the `-v` option:

```
$ docker run -d -v $(pwd)/postgres/:/docker-entrypoint-initdb.d/ -e POSTGRES_PASSWORD=secret postgres:latest
```

> **_NOTE_** The `-e` option specifies an environment variable, in this case the password for the privileged user on the database which is required by the container to run. 

And if we connect on the newly started container we can see that the new database and user exists:

```
root@4a7e441efd54:/# psql workshop student
psql (14.4 (Debian 14.4-1.pgdg110+1))
Type "help" for help.

workshop=>
workshop=> \conninfo
You are connected to database "workshop" as user "student" via socket in "/var/run/postgresql" at port "5432".
```

### Docker Volumes

We've seen how bind mounts are working and that are relying on the host filesystem. Docker volumes are a little bit different and are native to Docker, data being stored on a `special area` managed by docker. In order to create one we can use `docker volume`:

```
$ docker volume create workshop
```

Than you can visualize it using `docker volume ls`:

```
$ docker volume ls
DRIVER    VOLUME NAME
local     workshop
```

Also a volume can be created without a name, case in which it will be assigned a unique id:

```
$ docker volume create && docker volume ls
5dd3898c6ef60623c7ff24aa4f179f5c39cdcd5b418a728679955e2a2a428863
DRIVER    VOLUME NAME
local     5dd3898c6ef60623c7ff24aa4f179f5c39cdcd5b418a728679955e2a2a428863
local     workshop
```

In order to find out more about the newly created volume and see were is mounted we can use the inspect command:

```
$ docker volume inspect workshop
[
    {
        "CreatedAt": "2022-07-27T06:10:33Z",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/workshop/_data",
        "Name": "workshop",
        "Options": {},
        "Scope": "local"
    }
]
```

And in order to mount it on a container we have to use `-v`. Another thing to mention is that we can specify permissions when we are mounting it. E.g:

```
$ docker run -d -v workshop:/var/opt/project ubuntu:22.04
```

> **_NOTE_** You can specify permissions on the mounted volumes so that the container will have read-only or read-write.

## Docker Networking

The Docker networking subsystem is pluggable and uses drivers. Several such drivers exist by default, providing basic functionality for the network component. The default network driver is **bridge**, and it involves creating a software bridge that allows containers connected to the same network of this type to communicate with each other, while providing isolation from containers that are not connected to this bridge network. The Docker bridge driver automatically installs rules on the host machine so that containers on different bridge networks cannot communicate directly with each other. Bridge networks only apply to containers running on the same Docker machine.

For communication between containers running on different Docker machines, routing can be handled at the operating system level, or an overlay network can be used. As will be detailed later, overlay networks connect multiple Docker machines and allow services in a swarm to communicate with each other. Overlay networks can also be used to facilitate communication between a swarm service and a standalone container, or between two containers running on different Docker machines.

Other Docker network drivers are **host** (for standalone containers, removing the network isolation between the container and the Docker host, thus using the host's network infrastructure directly), **macvlan** (allows MAC addresses to be assigned to a container, making it appear as a physical device on the network), or **none**.

Containers on the same network can communicate without exposing ports via named DNS. This means that we can access a container not by its IP, but by its name. Ports must be exposed for communication with the outside world (host, off-network containers, etc.).

To demonstrate how bridged networking works in Docker, we'll first start two containers running Alpine. By default, any newly created Docker container will be on a network called "bridge", so in order to demonstrate that two containers that are not on the same network cannot communicate, it will first need to remove them from that network.

```
$ docker container run --name c1 -d -it alpine
 
f5a8653a325e8092151614d5a6a80b04b9410ea8b8a5fcfc4028f1ad33239ad9
```

```
$ docker container run --name c2 -d -it alpine

b063ad1ef7bd0ae82a7385582415e78938f7df531cef9eefc33e065af09cf92c
```

```
$ docker network disconnect bridge c1
```
```
$ docker network disconnect bridge c2
```

> **_NOTE_** In the docker run command above, the `--name` parameter gives the container a name (or alias) by which we can refer to it more easily.

At this moment, containers c1 and c2 are not part of any network. Next, we will try to ping from one container to another.

```
$ docker exec -it c1 ash
 
/ # ifconfig
lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)
 
/ # ping c2
ping: bad address 'c2'
 
/ # exit
```

It can be seen above that c1 only has a loopback IP address and cannot access c2. It would also be interesting to note that we attached to the container using the `docker exec` command by running a shell (ash is the shell on Alpine).

To create a bridged network in Docker, we can use the following command:
```
$ docker network create -d bridge c1-c2-bridge
 
8644b8accd2a14d10c9911c36635ca6b161449b3aa527db878a727ec1bf980d0
```

Further, we can view existing networks like this:

```
$ docker network ls
 
NETWORK ID          NAME                DRIVER              SCOPE
ecd72738aa59        bridge              bridge              local
8644b8accd2a        c1-c2-bridge        bridge              local
615363cafefa        host                host                local
1e3b8e49b20d        none                null                local
```

We can add a container to a network either when we start the container or when it has already been started. For the case above, where c1 and c2 were already on, we can add them to the network like this:

```
$ docker network connect c1-c2-bridge c1
$ docker network connect c1-c2-bridge c2
```

If c1 and c2 weren't already started, we could start them already attached to the c1-c2-bridge network like this:

```
$ docker container run --name c2 -d -it --network=c1-c2-bridge alpine
 
67dde5da9b793de63903ac85ff46574da77f0031df9b49acf44d58062687729c
```

```
$ docker container run --name c1 -d -it --network=c1-c2-bridge alpine
 
4de3e000700f81d31e0458dbd034abe90dfce6b1b992d23d760a44f748c0de0d
```

We can see the containers in a network like this:
```
$ docker network inspect c1-c2-bridge
 
[...]
"Containers": {
    "b063ad1ef7bd0ae82a7385582415e78938f7df531cef9eefc33e065af09cf92c": {
        "Name": "c2",
        "EndpointID": "a76463662d110804205e9211537e541eb0de2646fa90e8760d3419a6dc7d32c7",
        "MacAddress": "02:42:ac:12:00:03",
        "IPv4Address": "172.18.0.3/16",
        "IPv6Address": ""
    },
    "f5a8653a325e8092151614d5a6a80b04b9410ea8b8a5fcfc4028f1ad33239ad9": {
        "Name": "c1",
        "EndpointID": "95d9061b47f73f9b4cc7a82111924804bdc73d0b496549dec834216ee58c64ed",
        "MacAddress": "02:42:ac:12:00:02",
        "IPv4Address": "172.18.0.2/16",
        "IPv6Address": ""
    }
}
[...]
```

At this point, the two containers are part of the same network and can communicate:
```
$ docker exec -it c1 ash

/ # ping -c2 c2
PING c2 (172.18.0.3): 56 data bytes
64 bytes from 172.18.0.3: seq=0 ttl=64 time=6.258 ms
64 bytes from 172.18.0.3: seq=1 ttl=64 time=0.109 ms

--- c2 ping statistics ---
2 packets transmitted, 2 packets received, 0% packet loss
round-trip min/avg/max = 0.109/3.183/6.258 ms

/ # exit
```

```
$ docker exec -it c2 ash

/ # ping -c2 c1
PING c1 (172.18.0.2): 56 data bytes
64 bytes from 172.18.0.2: seq=0 ttl=64 time=0.111 ms
64 bytes from 172.18.0.2: seq=1 ttl=64 time=0.268 ms

--- c1 ping statistics ---
2 packets transmitted, 2 packets received, 0% packet loss
round-trip min/avg/max = 0.111/0.189/0.268 ms

/ # exit
```

## Docker Compose

Docker Compose is a utility created by Docker used to centralize the runtime configuration of containers in a declarative manner. Using Yet Another Markup Language (YAML) configuration files, Docker Compose centralizes the configuration process in a natural, declarative manner.

### Key Concepts

#### YAML file format

YAML files are typically used to write configurations declaratively. The format is very easy to understand and use, like this:

- "key:value" elements are used
- indented lines represent child properties of previous paragraphs
- the lists are delimited by "-".

##### Docker Compose example file

```
# docker-compose.yml
version: "3.8"
services:
    api:
        build: . # build image using a Dockerfile
        image: register-image-name:version # use an image from the current register
        environment:
            NODE_ENV: development
            ENVIRONMENT_VARIABLE: value
        ports:
            - "5000:80"
        networks:
            - network-docker
    
    postgres:
        image: postgres:12
        secrets:
            - my-secret-password
        environment:
            PGPASSWORD_FILE: /run/secrets/my-secret-password
        volumes:
            - volume-docker:/var/lib/postgresql/data
            - ./startup-scripts/init-db.sql:/docker-entrypoint-init.d/init-db.sql
        networks:
            - network-docker

volumes:
    volume-docker:

networks:
    network-docker:

secrets:
    my-secret-password:
        file: './my-not-so-secret-password.txt'
```

### Version

> **_NOTE_** It is mandatory to pass the version in any Docker Compose file.

#### Services

Describes the services/containers that will run after the configuration is started by Compose. Each service represents a container that will have the name and configuration of the service. In the example above, the containers will be called api and postgres. The most important properties of services are the following:

- **build** - specifies the directory where the Dockerfile from which to build the container is located

- **image** - specifies the name of the image used to run the container

- **ports** - a list of entries of type "host_port:service_port" where port exposure and mapping is performed

- **volumes** - a list of "host_volume:service_path" type entries where volume mappings are specified; the same rules that apply to the classic run are maintained here; "host_volume" can be a standard volume or a bind mount

- **networks** - the list of networks to which the service/container belongs

- **secrets** - the list of secrets that will be used within the service/container

- **environment** - object with entries of the type "service_environment_variable_name:value" that injects the specified environment variables when the service/container is running.

> **_NOTE_** The build and image options are disjoint.

> **_NOTE_** Secrets must also be passed within environment variables, according to the documentation. For example, in the Postgres configuration, secrets must be passed in special environment variables, suffixed with _FILE, along with their full path (ie /run/secrets/SECRET_NAME).

#### Volumes
Describes the volumes used in the configuration. Volumes are passed as objects. If the default configuration is not desired, the value is an empty field. 

> **_NOTE_** The top-level volumes property must be written at the same indentation level as services. Not to be confused with the child volumes property inside the services configuration.

#### Networks

Describes the networks used in the configuration. Networks are passed as objects. If the default configuration is not desired, the value is an empty field. An example network configuration is the following (where we use a network that already exists, because it was created independently of the Docker Compose file):

```
networks:
   my-already-existing-network:
        external: true
        name: original-already-existing-network
```
In the case above, my-already-existing-network is just a "rename" of an already existing network.

> **_NOTE_** The top-level networks property must be written at the same indentation level as services. Not to be confused with the networks child property inside the services configuration.

#### Secrets

Describes the secrets used in the configuration. They hold sensitive information in a secure, encrypted manner within Swarm, which we'll talk about in Lab 3. In Compose, secrets are not secure, but were introduced to ease the transition to Swarm. In Docker Compose, secrets can only come from external files, which must be specified for each secret.

The top-level secrets property is written at the same indentation level as services. Not to be confused with the secrets child property inside the services configuration.

#### Docker Compose Commands

```
$ docker compose start                          # start V2 containers
$ docker compose pause                          # pause a service's containers (send SIGPAUSE) V2
$ docker compose unpause                        # unpause V2 containers
$ docker compose ps                             # list active V2 containers
$ docker compose ls                             # list all V2 container stacks
$ docker compose -p my-project -f my-docker-compose.yml up # use the specified Compose file instead of the default one and with a V2 project name
$ docker compose down                           # stop containers and delete them, along with networks, volumes and images created at up V2
$ docker compose rm                             # delete all stopped containers (the name of the container to be deleted can also be specified at the end) V2
$ docker compose rm -s -v                       # -s stops all containers and -v also deletes V2 attached anonymous volumes
```

## Monitoring My First Application

When writing code, releasing a product another important part of your works is the monitoring stack, this is usually helping you identify various problems when your service is not booting due to some errors (e.g. panics, sigsegvs) or is underperforming
maybe it is underprovisioned in terms of ram or cpu (e.g. your brand new changes have changed the memory profile and so we might need to increase the limits for our hardware resources).

> **_NOTE_** In most of the cases the monitoring component is combined with an `alerting system` that can allow the service owner to take a proper action to mitigate the issue (e.g. increase resources, scale the service or create a patch for a code bug).

One way to investigate such problems is to use `docker logs` and `docker stats`. E.g:

```
$ docker run -d -v $(pwd)/postgres/:/docker-entrypoint-initdb.d/ -e POSTGRES_PASSWORD=secret postgres:latest
b2d3fe1510a49805c0a4c44efd291b074c6444ededbb0baeb1fa2d980b0dd225
$ docker ps
CONTAINER ID   IMAGE             COMMAND                  CREATED              STATUS              PORTS       NAMES
b2d3fe1510a4   postgres:latest   "docker-entrypoint.s…"   3 seconds ago        Up 2 seconds        5432/tcp    optimistic_gould
$ docker logs optimistic_gould
The files belonging to this database system will be owned by user "postgres".
This user must also own the server process.

The database cluster will be initialized with locale "en_US.utf8".
The default database encoding has accordingly been set to "UTF8".
The default text search configuration will be set to "english".

Data page checksums are disabled.

fixing permissions on existing directory /var/lib/postgresql/data ... ok
creating subdirectories ... ok
selecting dynamic shared memory implementation ... posix
selecting default max_connections ... 100
selecting default shared_buffers ... 128MB
selecting default time zone ... Etc/UTC
creating configuration files ... ok
running bootstrap script ... ok
performing post-bootstrap initialization ... ok
syncing data to disk ... ok
% docker stats optimistic_gould
CONTAINER ID   NAME               CPU %     MEM USAGE / LIMIT     MEM %     NET I/O     BLOCK I/O         PIDS
b2d3fe1510a4   optimistic_gould   0.05%     56.78MiB / 7.675GiB   0.72%     796B / 0B   47.3MB / 52.4MB   7
```

Unfortunately those commands are not providing a perspective in time of how the metrics are evolving therefore for the next part of the workshop we will focus on using more advanced systems like `prometheus` and `grafana`.

### Prometheus.io

Prometheus is basically a tool written in `golang` that was meant for monitoring. It is scraping at fixed intervals multiple HTTP endpoints collecting key-value metrics which are aggregated in a multi-dimensional data model. The scraped targets can be discovered via service discovery (e.g. using Consul or DNS). And it also comes with its own query language which is called PromQL.

> **_NOTE_** It has multiple components like `Push Gateway`, `Alertmanager`, `Exporters`, etc. But for our use case today we will only focus on how we configure the `Prometheus Server` since it defines a series of jobs that are run in order to collect the metrics. In `monitoring/prometheus` you can find an example of a file that configures a series of jobs that will scrape different targets for metrics.

### Grafana

Same as prometheus, grafana is an open-source tool that is generally used for running data analytics, pulling up metrics that make sense of the massive amount of data & to monitor our apps. Grafana connects with every possible data source, commonly referred to as databases such as Graphite, Prometheus, Influx DB, ElasticSearch, MySQL, PostgreSQL etc.

> **_NOTE_** For our use case today we will be mainly focusing on Prometheus and see how we can monitor multiple containers while creating some dashboards for docker metrics.

### cAdvisor

In order to get some data in grafana we need an `exporter` for prometheus. Such an alternative is `cAdvisor` which provides container users an understanding of the resource usage and performance characteristics of their running containers. It is a running daemon that collects aggregates processes and exports information about running containers. Specifically for each container it keeps resource isolation parameters, historical resource usage histograms of complete historical resource usage and network statistics.

First of all let's use the following docker-compose to deploy the aforementioned components:

```
version: "3.8"
services:
  
  prometheus:
    prometheus:
    image: prom/prometheus:latest
    container_name: prometheus
    volumes:
      - ./prometheus:/etc/prometheus
      - prometheus_data:/prometheus
    command:
      - "--config.file=/etc/prometheus/prometheus.yml"
      - "--storage.tsdb.path=/prometheus"
      - "--web.console.libraries=/etc/prometheus/console_libraries"
      - "--web.console.templates=/etc/prometheus/consoles"
      - "--storage.tsdb.retention.time=200h"
      - "--web.enable-lifecycle"
    restart: unless-stopped
    ports:
      - 9090:9090
    networks:
      - monitor-net
    depends_on:
      - cadvisor

  grafana:
    image: grafana/grafana:latest
    container_name: grafana
    volumes:
      - grafana_data:/var/lib/grafana
      - ./grafana/provisioning/dashboards:/etc/grafana/provisioning/dashboards
      - ./grafana/provisioning/datasources:/etc/grafana/provisioning/datasources
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=admin
      - GF_USERS_ALLOW_SIGN_UP=false
    restart: unless-stopped
    ports:
      - 3000:3000
    restart: unless-stopped
    networks:
      - monitor-net
    depends_on:
      - prometheus
  
  cadvisor:
    image: gcr.io/cadvisor/cadvisor:latest
    container_name: cadvisor
    ports:
      - 8080:8080
    restart: unless-stopped
    privileged: true
    devices:
      - /dev/kmsg:/dev/kmsg
    volumes:
      - /:/rootfs:ro
      - /var/run/docker.sock:/var/run/docker.sock:rw
      - /sys:/sys:ro
      - /var/lib/docker/:/var/lib/docker:ro
    networks:
      - monitor-net

networks:
  monitor-net:
    driver: bridge

volumes:
    prometheus_data: {}
    grafana_data: {}
```

> **_NOTE_** Everything required for this monitoring stack is located in `./monitoring`. Run the docker-compose file and then you should be able to see all the components using `localhost:[9090|8080|3000]`. For grafana the user and pass are admin and admin initially you will be prompted to provide a new password and confirm it, you can reuse the old one.
