#syntax=docker/dockerfile:1
FROM golang:1.17-alpine
WORKDIR ./app
COPY ./ ./
RUN go mod download
#COPY *.go ./
RUN go build ./cmd/main.go
#EXPOSE 8080:8080
#CMD [ "./main" ]