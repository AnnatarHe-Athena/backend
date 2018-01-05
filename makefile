output_name = backend-exec
BUILD_TIME=`date +%FT%T%z`

protoc:
	protoc -I proto-src/ proto-src/*.proto --go_out=proto --plugin=grpc:proto

release:
	go build -tags release -o $(output_name)
	mv $(output_name) /tmp
	echo "to /tmp and send it to server"
	echo "build date: ", BUILD_TIME