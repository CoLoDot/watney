FROM golang:latest

COPY . /app/

WORKDIR /app/api

CMD ["go", "run", "main.go"]
