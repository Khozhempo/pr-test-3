@cls
@cd ../cmd/pr-test-3
@..\..\third_party\protoc\bin\protoc.exe users/delivery/users.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.
@pause