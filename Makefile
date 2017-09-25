all: test readme

test:
	go vet ./...
	go test ./...

bench:
	go test -bench=. -benchmem

readme:
	godocdown 4d63.com/uuid > README.md

setup:
	go get github.com/robertkrimen/godocdown/godocdown
