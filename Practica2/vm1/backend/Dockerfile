FROM golang

WORKDIR /backend_app

COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]