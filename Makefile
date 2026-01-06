.PHONY: default build install clean

default: build install

build:
	cd core && go build && cd ..
	go build

install:
	sudo cp ./fountain /usr/local/bin/ff
	sudo cp ./fountain /usr/local/bin/fountain

clean:
	rm -rf fountain
	rm -rf core/f1

