FROM golang:latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /move

EXPOSE 8080

CMD ["/move"]





