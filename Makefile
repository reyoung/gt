all: bin/gt-server bin/gt-server.darwin.x64 bin/gt-server.darwin.arm64 bin/gt-client\
 	bin/gt-client.darwin.x64 bin/gt-client.darwin.arm64

clean:
	rm -rf bin/*

bin/gt-server:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gt-server ./cmd/server

bin/gt-server.darwin.x64:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/gt-server.darwin.x64 ./cmd/server

bin/gt-server.darwin.arm64:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/gt-server.darwin.arm64 ./cmd/server

bin/gt-client:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gt-client ./cmd/client

bin/gt-client.darwin.x64:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/gt-client.darwin.x64 ./cmd/client

bin/gt-client.darwin.arm64:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/gt-client.darwin.arm64 ./cmd/client
