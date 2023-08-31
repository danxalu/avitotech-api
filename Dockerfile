FROM golang:1.17-alpine3.16 AS starter
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
COPY . .
RUN go build -o /server cmd/api/main.go

FROM alpine:3.16
WORKDIR .
COPY cfg.yml .
COPY --from=starter /server /server
CMD ["/server"]