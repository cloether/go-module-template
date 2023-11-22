FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY  . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-module-template
EXPOSE 8080
CMD ["/go-module-template"]