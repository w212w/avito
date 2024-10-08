FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN go build -o tender_service .

EXPOSE 8080

CMD ["./tender_service"]
