# # # # # # # # # # # # # # # #
# Golang build
# # # # # # # # # # # # # # # # 
FROM golang:1.16
ENV FLAGS="-X 'helper.Builder=docker' -X 'helper.BuildTime=$(date)' -X 'helper.BuildVersion=${CI_JOB_ID}' '-extldflags=-static'"
WORKDIR /go/src/github.com/DarioCalovic/secretify
ADD . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -ldflags "${FLAGS}" -o secretify-api ./cmd/api.go
RUN chmod +x secretify-api

# # # # # # # # # # # # # # # #
# Node
# # # # # # # # # # # # # # # # 
FROM node:14-alpine

# Install packages
RUN apk add --update --no-cache supervisor

# Create app directory
WORKDIR /app

COPY --from=0 /go/src/github.com/DarioCalovic/secretify/secretify-api .
COPY supervisord.conf /etc/supervisord.conf

ADD ./ui/ /app/ui

RUN mkdir -p /app/db

# global install & update
RUN npm i -g npm
RUN npm i -g nuxt@2.15.4

RUN npm install /app/ui
RUN npm run build --prefix /app/ui

ENV HOST 0.0.0.0
EXPOSE 3000 8800

# start command
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
