
APPLICATION ?= weather-app

setup:
	mkdir -p build/bin

release:
	go build -o build/bin/${APPLICATION} main.go

install: release

run: release
	./build/bin/${APPLICATION}

clean:
	rm -rf build
