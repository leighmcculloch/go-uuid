test:
	go test -v -run=. -bench=. -benchmem

build:
	go install
