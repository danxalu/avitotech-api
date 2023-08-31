FROM golang:1.16 AS starter

WORKDIR /avitotech-api

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download
RUN go build -o /server cmd/api/main.go

COPY . .

FROM alpine:3.18
WORKDIR .

COPY config.yml ./

COPY --from=starter /server /server

CMD ["/server"]