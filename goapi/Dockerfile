FROM golang:1.19.4-alpine

EXPOSE 9090

ENV HOME=/home/louissantucci

WORKDIR /goapi

RUN apk add --no-cache gcc g++ git openssh-client

COPY ./goapi/go.mod ./
COPY ./goapi/go.sum ./

RUN go mod download

COPY ./goapi .

# Overwriting Docker config file
COPY ./docker/config/config.go ./config

# WARNING: environment dependency
# Copying SSL certificates to enable HTTPS: need to be changed for your usage
COPY $HOME/Documents/certs/certs/nginx.crt ./certs
COPY $HOME/Documents/certs/private/nginx.key ./certs

RUN go build -v -o goapi .

CMD [ "./goapi" ]


