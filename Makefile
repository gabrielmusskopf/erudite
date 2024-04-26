server:
	go build -o bin/erudite-server *.go
	./bin/erudite-server

eructl:
	go build -o bin/eructl cmd/eructl.go
