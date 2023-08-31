FROM golang:1.17 AS starter
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
COPY . .
RUN go build -o /server cmd/api/main.go

FROM alpine:3.18
WORKDIR .
COPY config.yml ./
COPY --from=starter /server /server
CMD ["/server"]