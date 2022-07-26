build:
	GOARCH=wasm GOOS=js go build -o ./web/app.wasm ./server/main.go
	go build -o ./bin/todo ./server/main.go

run: build
	./bin/todo