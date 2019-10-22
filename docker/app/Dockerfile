FROM golang:1.13-alpine as build-env
WORKDIR /bass

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash mysql-client

COPY --from=build-env /bass/bin/app /bass/
COPY --from=build-env /bass/bin/migrate /bass/
COPY --from=build-env /bass/migrations /bass/migrations

COPY --from=build-env /bass/docker/app/bin /usr/local/bin/bass/
RUN chmod +x /usr/local/bin/bass/*

EXPOSE 8080
CMD ["/bass/app"]