run:
	go build .
	./adjutant

build-docker:
	GOOS=linux GOARCH=amd64 go build -o adjutant_linux_amd64 .
	docker build -t adjutant .
	docker tag adjutant:latest ${DOCKER_REGISTRY}/adjutant:latest
	docker push ${DOCKER_REGISTRY}/adjutant:latest
