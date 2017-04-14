all: test readme

test:
	go test ./...

bench:
	go test -bench=. -benchmem

readme:
	godocdown github.com/leighmcculloch/go-uuid > README.md

setup:
	go get github.com/davecheney/godoc2md
