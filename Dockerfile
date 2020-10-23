FROM golang:latest

WORKDIR /app

COPY . /app

EXPOSE 8000

RUN cd api

CMD ["go", "run", "main.go"]
