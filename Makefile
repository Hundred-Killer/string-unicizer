build-win:
	GOOS=windows GOARCH=386 go build -o ./build/uniqueizer.exe ./cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/uniqueizer ./cmd/main.go

run:
	go run ./cmd