executable=qcon-gtin
repository=10.10.10.103:5000
container=qcon-gtin
tag=0.0.18

build/$(executable): *.go
	mkdir -p build
	go build -o build/$(executable)

dist/$(executable): *.go
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/$(executable)

build/container: dist/$(executable)
	docker build --no-cache -t $(container) .
	mkdir -p build
	touch build/container

.PHONY: run
run: build/container
	docker run -p 8080:80 $(container)

.PHONY: release
release: build/container
	docker tag -f $(container) $(repository)/$(container):$(tag)
	docker push $(repository)/$(container):$(tag)

.PHONY: clean
clean:
	rm -rf build
	rm -rf dist

