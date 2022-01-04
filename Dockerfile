
# pull official base image
FROM node:16-alpine3.11 as build-node
RUN apk --no-cache --virtual build-dependencies add \
        python \
        make \
        g++
WORKDIR /app
COPY ./frontend .
RUN yarn install
RUN yarn build

# Build stage
FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Build stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY --from=build-node /app/build ./build
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY backend/db/migration ./migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]