.PHONY: build

bin=bluebell
win_bin=bluebell_win
linux_bin=bluebell_linux

all: check build

build:
	go build -o $(bin)

run:
	./$(bin)

cross:
	CGO_ENABLE=0 GOOS=windows go build -o $(win_bin)
	CGO_ENABLE=0 GOOS=linux go build -o$(linux_bin)

check:
	go fmt ./
	go vet ./

run:
	go build -o $(bin)
	./$(bin)

debug:
	go build -o $(bin)
	./$(bin) --model debug

