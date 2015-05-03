executable=qcon-gtin
account=10.10.10.103:5000
tag=qcon-gtin
release=0.0.14

build/$(executable): *.go
	mkdir -p build
	go build -o build/$(executable)

dist/$(executable): *.go
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/$(executable)

build/container: dist/$(executable)
	docker build --no-cache -t $(executable) .
	mkdir -p build
	touch build/container

.PHONY: run
run: build/container
	docker run -p 8080:80 $(tag)

.PHONY: release
release: build/container
	docker tag -f $(executable) $(account)/$(tag):$(release)
	docker push $(account)/$(tag):$(release)

.PHONY: clean
clean:
	rm -rf build
	rm -rf dist

