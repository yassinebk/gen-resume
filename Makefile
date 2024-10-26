
.PHONY: build install 

build: 
	go build -o genresume .

install: build 
	mv ./genresume /usr/bin/
