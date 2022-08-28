gen-ioc:
	iocli gen
gen-api:
	go generate ./...
run-web:
	go run cmd/server/web/server.go -conf cmd/server/web/conf
run-user:
	go run cmd/server/user/server.go -conf cmd/server/user/conf
run-all-in-one:
	go run cmd/server/all-in-one/server.go -conf cmd/server/all-in-one/conf
visual:
	go-callvis -nostd cmd/servers/user/main.go
protoc-user: api/proto/user/*.proto
	protoc -I./api/proto/user \
	--go_out=plugins=grpc:. \
	--experimental_allow_proto3_optional \
	api/proto/user/*.proto