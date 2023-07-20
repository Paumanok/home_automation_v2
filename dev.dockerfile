FROM golang:latest

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app
RUN mkdir "/build"
COPY . .
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go version
EXPOSE 8080
ENTRYPOINT CompileDaemon -build="go build -tags dev -o /build/app -buildvcs=false" -command="/build/app"
