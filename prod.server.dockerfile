FROM golang:1.24-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p deploy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./deploy/app main.go

FROM alpine:latest

WORKDIR /app

COPY ./docs ./docs
COPY --from=build /app/deploy/app /app/app

ENTRYPOINT [ "./app" ]