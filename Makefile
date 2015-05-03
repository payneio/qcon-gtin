executable=qcon-gtin
account=payneio
tag=qcon-gtin
release=0.0.11

build/$(executable): *.go
	mkdir -p build
	go build -o build/$(executable)

build/container: dist/$(executable)
	docker build --no-cache -t $(executable) .
	mkdir -p build
	touch build/container

dist/$(executable): *.go
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/$(executable)

.PHONY: run
run: build/container
	docker run -p 8080:80 $(tag)

.PHONY: release
release: build/container
	docker tag -f $(tag) $(account)/$(tag):$(release)
	docker push $(account)/$(tag):$(release)

.PHONY: clean
clean:
	rm -rf build
	rm -rf dist

