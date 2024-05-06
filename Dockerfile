FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./gif-engine ./cmd/main.go

EXPOSE 5000

CMD ["/app/gif-engine"]