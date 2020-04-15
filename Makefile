.PHONY: build start

build: build-version

build-version:
	docker build -t smartip .
start:
	docker run -d -p 7001:6001 smartip

