FROM golang:1.22.5-bookworm

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -v -o app ./cmd/app/main.go

CMD /app/app