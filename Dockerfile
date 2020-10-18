FROM golang:latest

WORKDIR /app

COPY . /app

EXPOSE 8000

CMD ["go", "run", "."]
