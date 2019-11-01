xxx:
	@echo "Please choose an optimal action."

# Protocol Buffers
generate_pb:
	@make generate_pb_server
	@make generate_pb_web

clean_pb:
	@make clean_pb_server
	@make clean_pb_web

## Server
generate_pb_server:
	@protoc --proto_path ./proto --go_out=$(GOPATH)/src proto/*.proto
	@protoc --proto_path ./proto --go_out=$(GOPATH)/src proto/structs/*.proto

clean_pb_server:
	@rm -rf ./server/view/pb/messages/*.pb.go
	@rm -rf ./server/view/pb/messages/structs/*.pb.go

## Web
generate_pb_web:
	@echo "Not Implemented Yet."

clean_pb_web:
	@echo "Not Implemented Yet."