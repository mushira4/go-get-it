test:
	go test -v go-get-it/action_test
	go test -v go-get-it/config_test

install:
	go clean
	make dependencies
	make test
	go install

dependencies:
	dep ensure

