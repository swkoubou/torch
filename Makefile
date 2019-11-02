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
	@protoc --proto_path ./proto --go_out=server/view/pb proto/*.proto
	@protoc --proto_path ./proto --go_out=server/view/pb proto/structs/*.proto
	@mv server/view/pb/github.com/swkoubou/torch/server/view/pb/messages server/view/pb
	@rm -rf server/view/pb/github.com

clean_pb_server:
	@rm -rf ./server/view/pb/messages/*

## Web
generate_pb_web:
	@pbjs -t static-module -w es6 -o web/proto/web.js proto/*.proto proto/structs/*.proto
	@pbts -o web/proto/web.d.ts web/proto/web.js

clean_pb_web:
	@rm -rf web/proto/*.js
	@rm -rf web/proto/*.d.ts
