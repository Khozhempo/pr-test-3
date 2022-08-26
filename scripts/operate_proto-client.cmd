@cls
@cd ../cmd/pr-test-3-client
@..\..\third_party\protoc\bin\protoc.exe proto/users.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.
@pause