build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/gowizz ./*.go

clean:
	rm -rf ./bin ./vendor
