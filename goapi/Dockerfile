FROM golang:1.19.4-alpine

WORKDIR /goapi

RUN apk add --no-cache gcc g++ git openssh-client gettext

COPY ./goapi/go.mod ./
COPY ./goapi/go.sum ./
RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod download

# WARNING: environment dependency
# Copying SSL certificates to enable HTTPS: need to be changed for your usage
ADD ./certs/certificate ./certs/nginx.key
ADD ./certs/nginx.crt ./certs/

COPY ./goapi .

# Copy env variables into container for build
COPY ./docker/.env ./

# Overwriting Docker config file
COPY ./docker/goapi/config.template.go ./config
COPY ./docker/goapi/main.template.go ./

RUN source .env && envsubst '$${IP}' < ./config/config.template.go > ./config/config.go && rm ./config/config.template.go
RUN source .env && envsubst '$${IP}' < ./main.template.go > ./main.go && rm main.template.go

# Generate swagger config
RUN swag init

RUN go build -v -o goapi .

EXPOSE 9090
ENV GIN_MODE=release
CMD [ "./goapi" ]