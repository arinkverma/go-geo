all: download build run
	

build:
	docker build . -t go-geo:latest

run:
	docker-compose up

download:
	rm -rf tmp
	mkdir -p tmp
	curl -o tmp/cities15000.zip https://download.geonames.org/export/dump/cities15000.zip
	unzip tmp/cities15000.zip cities15000.txt -d tmp