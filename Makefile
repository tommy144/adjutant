run:
	go build .
	./adjutant

build-docker:
	GOOS=linux GOARCH=amd64 go build -o adjutant_linux_amd64 . 
	docker build .
	docker build -t memesofourlives/adjutant:latest .

