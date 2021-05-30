FROM golang:1.16.4-buster

RUN mkdir /code
WORKDIR /code

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src /code/.
