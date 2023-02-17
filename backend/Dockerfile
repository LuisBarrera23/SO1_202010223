FROM golang

WORKDIR /backend_app

COPY . .

RUN go mod download

EXPOSE 5000

CMD ["go", "run", "main.go"]