test:
	go test -v go-get-it/action_test
	go test -v go-get-it/config_test

install:
	go clean
	make dependencies
	make test
	go install

dependencies:
	go get github.com/garyburd/redigo/redis
	go get github.com/ghodss/yaml


