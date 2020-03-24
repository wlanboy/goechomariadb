# goechomariadb
Golang rest service based on echo and gorm using mariadb

# build
* go get -d -v
* go clean
* go build

# depends on
* MariaDB instance: https://github.com/wlanboy/Dockerfiles/tree/master/MariaDB

# run
* go run main.go

# debug
* go get -u github.com/go-delve/delve/cmd/dlv
* dlv debug ./goechomariadb

# dockerize (docker image size is 9.89MB)
* GOOS=linux GOARCH=386 go build (386 needed for busybox)
* GOOS=linux GOARCH=arm GOARM=6 go build (Raspberry Pi build)
* GOOS=linux GOARCH=arm64 go build (Odroid C2 build)
* docker build -t goechomariadb .

# run docker container
*docker run -d -p 8000:8000 goechomariadb

# call
* curl -X POST http://127.0.0.1:8000/api/v1/event -H 'Content-Type: application/json' -d '{"name": "test", "type": "info"}'
* curl -X GET http://127.0.0.1:8000/api/v1/event 
