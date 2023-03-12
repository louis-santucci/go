FROM golang:1.19.4-alpine

EXPOSE 9090

WORKDIR /goapi

RUN apk add --no-cache gcc g++ git openssh-client

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o goapi

CMD [ "./goapi" ]


