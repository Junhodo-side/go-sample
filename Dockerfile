FROM golang:1.17.6-bullseye

# Change working dir -> for locating configrations
RUN mkdir -p /app
WORKDIR /app


COPY . .
RUN go mod download
RUN go install cmd/main.go
ENTRYPOINT [ "/go/bin/main" ]