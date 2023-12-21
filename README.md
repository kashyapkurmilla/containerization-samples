# Python and Go Application Containerization

This repository demonstrates the containerization of both Python and Go applications using Docker. It contains Dockerfiles and configuration files necessary to build and run Docker containers for these applications.

## Contents

- [Setup](#setup)
- [Python Application](#python-application)
- [Go Application](#go-application)
- [Running the Containers](#running-the-containers)
- [Contributing](#contributing)
- [License](#license)

## Setup

To use these Dockerfiles, ensure you have Docker installed on your system. You can download and install Docker from [Docker's official website](https://www.docker.com/get-started).

## Python Application

The `python-app` directory contains a sample Python application that demonstrates basic functionalities. The `Dockerfile` in this directory is responsible for building a Docker image for the Python application.

## Go Application

The `go-app` directory contains a sample Go application that showcases specific features. The `Dockerfile` in this directory is responsible for building a Docker image for the Go application.

## Running the Containers

### Python Application

To build and run the Python application container:

```bash
cd python-app
docker build -t python-app .
docker run -p 8000:8000 python-app
