# Madrid Recicla Server <!-- omit in toc -->

[![Linkedin](https://img.shields.io/badge/LinkedIn-carlosmartinezm-blue)](https://www.linkedin.com/in/carlosmartinezm/)

**Madrid Recicla Server** is part of the distributed application **Madrid Recicla**.

- [Madrid Recicla Database]
- [Madrid Recicla Server]
- [Madrid Recicla Web]

**Madrid Recicla Server** provides backend support to Madrid Recicla suite, allowing functionalities like:

- Geospatial listing of _recycling points of interest_* based on user proximity. 
- Manual load of static data provided by Comunidad de Madrid.
- Automatic load of data fetched from Comunidad de Madrid's public APIs.
  
>\* Recycling points of interest include fixed and mobile recycling points, and different containers for clothes, vegetal oil, batteries, paper and paperboard, glass, plastic, organic and others.

## Contents <!-- omit in toc -->
- [Before starting](#before-starting)
- [Consuming the service](#consuming-the-service)
- [Developer information](#developer-information)
  - [Environment variables](#environment-variables)
  - [Running the application in a Docker container](#running-the-application-in-a-docker-container)
  - [Running the application in your local machine](#running-the-application-in-your-local-machine)
- [Others](#others)

## Before starting

Before cloning this repository make sure you've read the [Madrid Recicla Development Template] guide. This guide will help you to setup your local environment.

## Consuming the service

Madrid Recicla server provides an external API to support frontend applications in tasks fetching geospatial data filtered by location and providing map configurations. It also provides an internal API for management tasks like loading static data related to recycling points of interest.

Check the [external API] and [internal API] definitions for more information.

## Developer information

### Environment variables

Madrid Recicla Server uses a set of environment variables that need to be setup before running the application. To do so, create a file called `.env` in the root directory of the project with proper values for the following keys:

```properties
PORT=${SERVER_PORT} # Server port
DB_CONNECTION_URI=${DB_CONNECTION_URI} # Database connection URI
DB_NAME=${DATABASE_NAME} # Database name
ALLOWED_ORIGINS=${ALLOWED_ORIGINS} # CORS origins
MAPBOX_TOKEN=${MAPBOX_TOKEN} # Mapbox token
```

`DB_CONNECTION_URI` uses different formats depending on the environment to connect.

- For local connection, the format is:
    ```sh
    DB_CONNECTION_URI=mongodb://${SERVER_DB_USER}:${SERVER_DB_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/?authMechanism=SCRAM-SHA-1&authSource=${DB_NAME}
    ```
- For remote connection, the format is:
    ```sh
    DB_CONNECTION_URI=mongodb+srv://${SERVER_DB_USER}:${SERVER_DB_PASSWORD}@${DB_HOST_PRO}/${DB_NAME}?retryWrites=true&w=majority
    ```

>⚠️ Notice that `.env` file is ignored for commits in `.gitignore` file as this should never be pushed to the repo.

### Running the application in a Docker container
The source code contains a `Dockerfile` and a `.dockerignore` file that allow for quick building and running a **Go** image in local [Docker Containers].

While located in the directory where these files are located, run the following commands:

To build the image

```zsh
% docker build -t madrid-recicla-server .
```
> \* Notice the dot (.) at the end of the command.

To run the image inside a Docker container

```zsh
% docker run -p 8080:8080 --env-file .env madrid-recicla-server
```
> \* `-p ${SERVER_PORT}:${SERVER_CONTAINER_PORT}` will expose port _SERVER_CONTAINER_PORT_ inside the container to port _SERVER_PORT_ outside the container.

### Running the application in your local machine

You can, however, manually run your server without using Docker by following these steps:

1. Modify `main` function in `src/main.go` to load your `.env` file:

    ```go
    package main

    import (
        "github.com/joho/godotenv"
        ...
    )

    func main() {
        err := godotenv.Load("path/to/your/.env")
        if err != nil {
            panic(err)
        }
        ...
    }
    ```

2. Upon importing the new dependency `"github.com/joho/godotenv"`, `go.mod` and `go.sum` files must be updated. To do so, locate in `src/`, where `main.go` file is, and run the following command:

    ```zsh
    % go mod tidy
    ```

3. Then, run the application:

    ```zsh
    % go run main.go
    ```

## Others

[![Linkedin](https://img.shields.io/badge/LinkedIn-carlosmartinezm-blue)](https://www.linkedin.com/in/carlosmartinezm/)
[![Hex.pm](https://img.shields.io/hexpm/l/plug)](http://www.apache.org/licenses/LICENSE-2.0)
[![Open Source](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://opensource.org/)

<!-- Links -->
[Madrid Recicla Database]: <https://github.com/martinezcarlos/madrid-recicla-dev-template/blob/main/db/README.md>
[Madrid Recicla Server]: <https://github.com/martinezcarlos/madrid-recicla-server>
[Madrid Recicla Web]: <https://github.com/martinezcarlos/madrid-recicla-web>
[Madrid Recicla Development Template]: <https://github.com/martinezcarlos/madrid-recicla-dev-template>
[Docker Containers]: <https://docs.docker.com/language/golang/>
[external API]: <http://localhost:8081>
[internal API]: <http://localhost:8082>
