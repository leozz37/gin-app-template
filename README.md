# Golang Gin template

🚀 This is a template for web apps using Golang, Gin, Docker and Kubernetes!

![cover](https://miro.medium.com/max/5000/1*HyemvyVt7JI25k-_cTKMcg.png)


[![Build](https://github.com/leozz37/gin-app-template/actions/workflows/build.yml/badge.svg)](https://github.com/leozz37/gin-app-template/actions/workflows/build.yml)
[![Unit Tests](https://github.com/leozz37/gin-app-template/actions/workflows/unit_tests.yml/badge.svg)](https://github.com/leozz37/gin-app-template/actions/workflows/unit_tests.yml)
[![Docker](https://github.com/leozz37/gin-app-template/actions/workflows/docker.yml/badge.svg)](https://github.com/leozz37/gin-app-template/actions/workflows/docker.yml)
[![Terraform](https://github.com/leozz37/gin-app-template/actions/workflows/terraform.yml/badge.svg)](https://github.com/leozz37/gin-app-template/actions/workflows/terraform.yml)

## Contents

- [Architecture](#architecture)
  - [Directory Structure](#directory-structure)
- [Building](#building)
  - [Binary](#binary)
  - [Makefile](#makefile)
  - [Docker](#docker)
  - [Docker-compose](#docker-compose)
  - [Kubernetes](#kubernetes)
- [Testing](#testing)
  - [Unit Tests](#unit-tests)
- [Continuous Integration](#continuous-integration)

### Architecture

This is an REST API, made with [Golang](https://golang.org/) and [Gin](https://github.com/gin-gonic/gin). You can manually run it or use [docker-compose](https://docs.docker.com/compose/install/) (recommended) or [Kubernetes](https://kubernetes.io/docs/setup/) to get everthing up.

### Directory Structure

The project is divided by the following strucure:

```txt
.github
    |__ workflows
k8s
routes
services
    |__ db
```

These following directory has some misc files, not directly related to the API code:

- [.github/](./github) holds the GitHub templates.

- [.github/worflows](.github/worflows) has the CI pipelines for GitHub Actions.

- [k8s](./k8s) Kubernetes services and pods config files.

For the API, this is the direcoty rules:

- [routes](./routes) has the Gin definitions of the endpoints routes.

- [services](./services) holds the services used across the code like the database.

- [services/db](./services/db) holds the databases integrations (MySQL and Redis).

## Building

By far, the easiest way to get everything running is with `docker-compose`. See the [docker-compose](#docker-compose) section. First, set up the environment variables for your project. You can user the `.env.example`:

```shell
$ source .env.example
```

### Binary

To build the binary, run the following:

```shell
$ go build -o ${APP_NAME}
```

To run the binary, run the following:

```shell
$ ./${APP_NAME}
```

Or simply:

```shell
$ go run main.go
```

### Makefile

To run the through the Makefile, run the following

```shell
$ source .env

$ make run
```

### Docker

Make sure you have [Docker](https://www.docker.com/get-started) installed on your machine.

To build the Docker image, run the following:

```shell
$ docker build . -t ${APP_NAME}            
```

To run the Docker image, run:

```shell
$ docker run -p $PORT:$PORT ${APP_NAME}
```

### docker-compose

To run the docker-compose:

```shell
$ docker-compose up
```

### Kubernetes

To create the Kubernetes pods, make sure you have [MinuKube](https://minikube.sigs.k8s.io/docs/start/) or [Kind](https://kind.sigs.k8s.io/) installed. In this example, we'll be using `kind`.

Then run the following command:

Create a cluster:

```shell
$ kind create cluster --name=${APP_NAME}
```

Select the cluster you just created:

```shell
$ kubectl cluster-info --context kind-${APP_NAME}
```

Create the application pods:

```shell
$ kubectl apply -f k8s/
```

Verify pods status:

```shell
$ kubectl get pods
```

Wait until all the pods STATUS are `running`.

You can expose the API service port, running:

```shell
$ kubectl port-forward service/${APP_NAME} 8080:8080
```

## Testing

The unit testes are written with the default testing tool of Golang.

### Unit Tests

To run the unit tests, do the following:

```shell
$ go test -v ./...
```

To run the tests with coverage, do the following:

```shell
$ go test -v -covermode=count ./...
```

## Continuous Integration

We use GitHub Actions for our CI tool. Right now we have four workflows, and you can check they state [here](https://github.com/leozz37/gin-app-template/actions):

[![Build](https://github.com/leozz37/gin-app-template/actions/workflows/build.yml/badge.svg)](https://github.com/leozz37/gin-app-template/actions/workflows/build.yml) - Building the binary status.

[![Unit Tests](https://github.com/leozz37/gin-app-template/actions/workflows/unit_tests.yml/badge.svg)](https://github.com/leozz37/gin-app-template/actions/workflows/unit_tests.yml) - Unit tests status.

[![Docker](https://github.com/leozz37/gin-app-template/actions/workflows/docker.yml/badge.svg)](https://github.com/leozz37/gin-app-template/actions/workflows/docker.yml) - Building the Docker image status.

[![Terraform](https://github.com/leozz37/gin-app-template/actions/workflows/terraform.yml/badge.svg)](https://github.com/leozz37/gin-app-template/actions/workflows/terraform.yml) - Terraform plan
