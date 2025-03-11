build:
	go build .

init:
	go mod download

extract-file:
	gunzip -c 10m-v2.txt.gz >10m-v2.txt

run-with-file:
	make extract-file
	make build
	./lottery-counter 10m-v2.txt

run:
	make build
	@echo using $(file)
	./lottery-counter $(file)
