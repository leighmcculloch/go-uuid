test:
	go test ./...

bench:
	go test ./... -bench=. -benchmem

build:
	go install
