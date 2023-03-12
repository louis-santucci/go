FROM golang:1.19.4-alpine

WORKDIR /goapi

RUN apk add --no-cache gcc g++ git openssh-client

COPY ./goapi/go.mod ./
COPY ./goapi/go.sum ./

RUN go mod download

# WARNING: environment dependency
# Copying SSL certificates to enable HTTPS: need to be changed for your usage
ADD ./certs/private/certificate ./certs/nginx.key
ADD ./certs/certs/nginx.crt ./certs/

COPY ./goapi .

# Overwriting Docker config file
COPY ./docker/config/config.go ./config

RUN go build -v -o goapi .

EXPOSE 9090

