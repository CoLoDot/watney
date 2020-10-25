all: up

build:
		docker build -t watney .

run:
		docker run -p 8080:8080 watney

test:
		cd api/ && go test ./... && go tool cover -func=coverage.out

test_coverage:
		cd api/ && go test ./... -coverprofile coverage.out && go tool cover -html=coverage.out
