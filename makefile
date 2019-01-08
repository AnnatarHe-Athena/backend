output_name = backend-exec
BUILD_TIME=`date +%FT%T%z`

protoc:
	# protoc -I proto-src/ proto-src/*.proto --go_out=proto --plugin=grpc:proto
	protoc -I proto-src/ proto-src/*.proto --go_out=plugins=grpc:proto

tests:
	go test ./...

dev:
	go run main.go
	# gin --port=9002 --appPort=9988 --notifications run main.go

release:
	CGO_ENABLED=0 go build -tags release -o $(output_name)
	mv $(output_name) /tmp
	echo "to /tmp and send it to server"
	echo "build date: ", BUILD_TIME