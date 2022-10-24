build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

nodemon:
	nodemon --exec go run main.go --signal SIGTERM