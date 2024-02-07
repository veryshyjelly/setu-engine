FROM golang:1.21.6

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /setu-engine

EXPOSE 8070

CMD ["/setu-engine"]