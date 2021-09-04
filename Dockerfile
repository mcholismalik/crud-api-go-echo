# build stage
FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .

# RUN go mod init eclaim-api
RUN go mod tidy
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init
# RUN apk add --no-cache git
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/api ./main.go

# final stage
FROM alpine:latest
# RUN apk add --no-cache git

WORKDIR /usr/app

COPY --from=build /go/src/app/bin /go/bin
COPY --from=build /go/src/app/ ./
COPY ./.env.development /
EXPOSE 3030
ENTRYPOINT ENV=DEV /go/bin/api