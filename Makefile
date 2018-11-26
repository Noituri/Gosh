NAME=Gosh
SOURCE=main.go

.PHONY: run
run:
	DEV=1 go run $(SOURCE)

.PHONY: build
build:
	go build -o release/Gosh $(SOURCE)

.PHONY: install
install: build
	sudo mv release/$(NAME) /usr/bin/$(NAME)