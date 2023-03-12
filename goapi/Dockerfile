FROM golang:1.19.4-alpine

EXPOSE 9090

WORKDIR /goapi

RUN apk add --no-cache gcc g++ git openssh-client

COPY ./goapi/go.mod ./
COPY ./goapi/go.sum ./

RUN go mod download

COPY ./goapi .

# Overwriting Docker config file
COPY ./docker/config/config.go ./config

RUN go build -v -o goapi .

CMD [ "./goapi" ]


