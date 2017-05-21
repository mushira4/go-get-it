install:
	go clean
	go get
	go test
	go install

dependencies:
	go get github.com/garyburd/redigo/redis
	go get github.com/ghodss/yaml


