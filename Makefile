build:
	go build -o ./bin/guardian .

install:
	sudo cp ./bin/guardian /usr/bin

