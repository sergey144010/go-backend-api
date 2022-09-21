FROM golang:latest

COPY ./data /data
WORKDIR /usr/src/app
COPY . .
RUN go mod download && go mod verify
RUN go build
CMD ["./go-backend-api"]