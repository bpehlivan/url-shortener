# Url Shortener
Simple url shortener project for practising go. Currenly not working as expected, errors to be fixed soon.

errors to be fixed:
- redirection not working as expected (redirects to internal path)

missing features(to be implemented):
- unit testing
- logging mechanism
- simple UI interface
- caching
- Setting expiry date for urls

## Setting Up The Project
Following steps are only for local development.

An environment file named `.env` should be defined in order to run the project inside a docker container. Sample variables can be found inside `local.env` file. Copying `local.env` file as `.env` is the shortest solution.

At the next step, creating a container image is required:
```shell
make build
```

## Running The Project
In order to run the project with docker compose file:
```shell
make up
```

By default, makefile uses `development.docker-compose.yml` file.

## Accessing Shell
```shell
make bash
```

After entering the interactive container shell, custom commands can be run. Following commands are starting with `go run main.go` because building binary steps are not included in the docker file.

## Available Commands
To access command-line options:
```shell
go run main.go help
```

Running migrations:
```shell
go run main.go migrate
```

Serving the web application:
```shell
go run main.go serve
```

## The API
### Creating Redirect Url

Method: `POST` | Endpoint: `<base-url>/urls/`
Request:
```json
{
  "endpoint": "www.google.com"
}
```

Sample Response:
```json
{
    "ID": "2f9e4d1c-e609-4e44-b87d-d6fa3e2e64bb",
    "CreatedAt": "2021-05-30T21:53:21.0367343Z",
    "UpdatedAt": "2021-05-30T21:53:21.0367343Z",
    "DeletedAt": null,
    "Url": "www.google.com",
    "KeyWord": "fpllngzi",
    "ValidUntil": "0001-01-01T00:00:00Z"
}
```

### Redirection
Method: `GET` | Endpoint: `<base-url>/urls/{endpointKeyword}`