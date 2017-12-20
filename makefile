
protoc:
	protoc -I proto/ proto/cells.proto --go_out=plugins=grpc:proto