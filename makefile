install:
	go clean
	go get
	go test
	go install

dependencies:
	go get github.com/garyburd/redigo/redis

