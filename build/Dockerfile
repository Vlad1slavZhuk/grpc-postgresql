FROM golang:alpine AS build
WORKDIR /app/
COPY . /app/
RUN CGO_ENABLED=0 go build -o /bin/demo